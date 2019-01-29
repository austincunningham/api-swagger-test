/*
 * Mobile Security Service Server API
 *
 * This is a sample mobile security service server.
 *
 * API version: 1.0.0
 * Contact: dffrench@redhat.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AppsInner struct {
	Id int64 `json:"id,omitempty"`
	AppId int64 `json:"appId,omitempty"`
	AppName string `json:"appName,omitempty"`
	NumOfDeployedVersions int64 `json:"numOfDeployedVersions,omitempty"`
	NumOfClients int64 `json:"numOfClients,omitempty"`
	NumOfAppLaunches int64 `json:"numOfAppLaunches,omitempty"`
}