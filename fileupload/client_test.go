package fileupload

import (
  "testing"

  stripe "github.com/stripe/stripe-go"
  . "github.com/stripe/stripe-go/utils"
)

const (
  expectedSize = 12412
  expectedMime = "application/pdf"
)

func init() {
  stripe.Key = GetTestKey()
}

func TestFileUploadNew(t *testing.T) {
  f, err := os.Open("/home/wangjohn/Documents/PebbleBeach.pdf")
  if err != nil {
    t.Errorf("Unable to open test file upload file %v\n", err)
  }

  uploadParams := &stripe.FileUploadParams{
    Purpose: "dispute_evidence",
    File: os.Open("/home/wangjohn/Documents/PebbleBeach.pdf"), // TODO(wangjohn): turn this into a real test pdf.
  }
  target, err := New(uploadParams)
  if err != nil {
    t.Error(err)
  }

  if target.Size != expectedSize {
    t.Errorf("Size %v does not match expected size %v\n", target.Size, expectedSize)
  }

  if target.Purpose != uploadParams.Purpose {
    t.Errorf("Purpose %v does not match expected purpose %v\n", target.Purpose, uploadParams.Purpose)
  }

  if target.Mime != expectedMime {
    t.Errorf("Mime %v does not match expected mime %v\n", target.Mime, expectedMime)
  }
}
