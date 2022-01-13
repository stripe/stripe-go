//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"bytes"
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
)

// The [purpose](https://stripe.com/docs/file-upload#uploading-a-file) of the uploaded file.
type FilePurpose string

// List of values that FilePurpose can take
const (
	FilePurposeAccountRequirement               FilePurpose = "account_requirement"
	FilePurposeAdditionalVerification           FilePurpose = "additional_verification"
	FilePurposeBusinessIcon                     FilePurpose = "business_icon"
	FilePurposeBusinessLogo                     FilePurpose = "business_logo"
	FilePurposeCustomerSignature                FilePurpose = "customer_signature"
	FilePurposeDisputeEvidence                  FilePurpose = "dispute_evidence"
	FilePurposeDocumentProviderIdentityDocument FilePurpose = "document_provider_identity_document"
	FilePurposeFinanceReportRun                 FilePurpose = "finance_report_run"
	FilePurposeFoundersStockDocument            FilePurpose = "founders_stock_document"
	FilePurposeIdentityDocument                 FilePurpose = "identity_document"
	FilePurposeIdentityDocumentDownloadable     FilePurpose = "identity_document_downloadable"
	FilePurposePCIDocument                      FilePurpose = "pci_document"
	FilePurposeSelfie                           FilePurpose = "selfie"
	FilePurposeSigmaScheduledQuery              FilePurpose = "sigma_scheduled_query"
	FilePurposeTaxDocumentUserUpload            FilePurpose = "tax_document_user_upload"
)

// Returns a list of the files that your account has access to. The files are returned sorted by creation date, with the most recently created files appearing first.
type FileListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Purpose      *string           `form:"purpose"`
}

// FileFileLinkDataParams is the set of parameters allowed for the
// file_link_data hash.
type FileFileLinkDataParams struct {
	Params    `form:"*"`
	Create    *bool  `form:"create"`
	ExpiresAt *int64 `form:"expires_at"`
}

// To upload a file to Stripe, you'll need to send a request of type multipart/form-data. The request should contain the file you would like to upload, as well as the parameters for creating a file.
//
// All of Stripe's officially supported Client libraries should have support for sending multipart/form-data.
type FileParams struct {
	Params `form:"*"`

	// FileReader is a reader with the contents of the file that should be uploaded.
	FileReader io.Reader

	// Filename is just the name of the file without path information.
	Filename     *string
	Purpose      *string
	FileLinkData *FileFileLinkDataParams
}

// This is an object representing a file hosted on Stripe's servers. The
// file may have been uploaded by yourself using the [create file](https://stripe.com/docs/api#create_file)
// request (for example, when uploading dispute evidence) or it may have
// been created by Stripe (for example, the results of a [Sigma scheduled
// query](https://stripe.com/docs/api#scheduled_queries)).
//
// Related guide: [File Upload Guide](https://stripe.com/docs/file-upload).
type File struct {
	APIResource
	Created   int64         `json:"created"`
	ExpiresAt int64         `json:"expires_at"`
	Filename  string        `json:"filename"`
	ID        string        `json:"id"`
	Links     *FileLinkList `json:"links"`
	Object    string        `json:"object"`
	Purpose   FilePurpose   `json:"purpose"`
	Size      int64         `json:"size"`
	Title     string        `json:"title"`
	Type      string        `json:"type"`
	URL       string        `json:"url"`
}

// FileList is a list of Files as retrieved from a list endpoint.
type FileList struct {
	APIResource
	ListMeta
	Data []*File `json:"data"`
}

// GetBody gets an appropriate multipart form payload to use in a request body
// to create a new file.
func (f *FileParams) GetBody() (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if f.Purpose != nil {
		err := writer.WriteField("purpose", StringValue(f.Purpose))
		if err != nil {
			return nil, "", err
		}
	}

	if f.FileReader != nil && f.Filename != nil {
		part, err := writer.CreateFormFile(
			"file",
			filepath.Base(StringValue(f.Filename)),
		)

		if err != nil {
			return nil, "", err
		}

		_, err = io.Copy(part, f.FileReader)
		if err != nil {
			return nil, "", err
		}
	}

	if f.FileLinkData != nil {
		values := &form.Values{}
		form.AppendToPrefixed(values, f.FileLinkData, []string{"file_link_data"})

		params, err := url.ParseQuery(values.Encode())
		if err != nil {
			return nil, "", err
		}
		for key, values := range params {
			err := writer.WriteField(key, values[0])
			if err != nil {
				return nil, "", err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.Boundary(), nil
}

// UnmarshalJSON handles deserialization of a File.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *File) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type file File
	var v file
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = File(v)
	return nil
}
