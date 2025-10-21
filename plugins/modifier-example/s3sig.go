package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"
)

var s3Expires = "604800"

func getSigV4Elements(req RequestWrapper) (string, string, string, error) {
	headers := req.Headers()
	queries := req.Query()

	access := "access"
	scope := "20251023/us-east-1/s3/aws4_request"
	signingKey := base64.StdEncoding.EncodeToString(getSigningKey())
	canonicalDate := "20251023T150405Z"

	credential := fmt.Sprintf("%s/%s", access, scope)

	delete(headers, "X-Amz-Date")
	delete(headers, "X-Amz-Content-Sha256")

	queries.Add("X-Amz-Algorithm", "AWS4-HMAC-SHA256")
	queries.Add("X-Amz-Credential", credential)
	queries.Add("X-Amz-Date", canonicalDate)
	queries.Add("X-Amz-Expires", s3Expires)

	canonicalHeaders := ""

	headerKeys := slices.Collect(maps.Keys(headers))
	slices.Sort(headerKeys)
	signedHeaders := strings.ToLower(strings.Join(headerKeys, ";"))

	for _, key := range headerKeys {
		if len(headers[key]) > 0 && headers[key][0] != "" {
			canonicalHeaders += strings.ToLower(key) + ":" + headers[key][0] + "\n"
		}
	}

	canonicalContent := "UNSIGNED-PAYLOAD"
	queries.Add("X-Amz-SignedHeaders", signedHeaders)
	req.URL().RawQuery = queries.Encode()

	canonicalRequest := strings.Join([]string{
		req.Method(),
		req.Path(),
		queries.Encode(),
		canonicalHeaders,
		signedHeaders,
		canonicalContent,
	}, "\n")

	// Sign the canonical request
	signature, err := calcS3Signature(canonicalRequest, scope, canonicalDate, signingKey)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to calculate signature for the request: %s", err.Error())
	}

	return signature, credential, signedHeaders, nil
}

func getSigningKey() []byte {
	t := time.Now()
	awsDate := t.Format("20060102")
	hmacDate := hmacsha256([]byte("AWS4secret"), awsDate)
	hmacRegion := hmacsha256(hmacDate, "us-east-1")
	hmacService := hmacsha256(hmacRegion, "s3")

	return hmacsha256(hmacService, "aws4_request")
}

func calcS3Signature(request, scope, timestamp, signingKey string) (string, error) {
	hashedRequest := sha256.Sum256([]byte(request))
	signatureBase := strings.Join([]string{
		"AWS4-HMAC-SHA256",
		timestamp,
		scope,
		hex.EncodeToString(hashedRequest[:]),
	}, "\n")

	keySlice, err := base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		return "", fmt.Errorf("could not decode signing key: %w", err)
	}

	signer := hmacsha256(keySlice, signatureBase)

	return hex.EncodeToString(signer), nil
}

func hmacsha256(key []byte, data string) []byte {
	hash := hmac.New(sha256.New, key)
	_, _ = hash.Write([]byte(data))

	return hash.Sum(nil)
}
