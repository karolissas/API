/*
 * Keep Charged
 *
 * API documentation
 *
 * API version: 0.0.1 beta
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse404 struct {
	// Operation status
	Success bool `json:"success,omitempty"`

	Errors []DefaultError `json:"errors,omitempty"`
}
