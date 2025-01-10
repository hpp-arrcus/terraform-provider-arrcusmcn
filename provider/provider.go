package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	b64 "encoding/base64"

	datacredential "github.com/Arrcus/terraform-provider-arrcusmcn/datasource/credential"
	datadeployment "github.com/Arrcus/terraform-provider-arrcusmcn/datasource/deployment"
	arc_resource "github.com/Arrcus/terraform-provider-arrcusmcn/resource"
	credentialresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/credential"
	deploymentresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/deployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &ArcOrchProvider{}

type ArcOrchProvider struct {
	Version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ArcOrchProvider{
			Version: version,
		}
	}
}

type ArcOrchProviderModel struct {
	ApiKey   types.String `tfsdk:"api_token"`
	Endpoint types.String `tfsdk:"endpoint"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Port     types.String `tfsdk:"port"`
}

func (p *ArcOrchProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "arrcusmcn"
	resp.Version = p.Version
}

func (p *ArcOrchProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "user name",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "password",
				Optional:            true,
				Sensitive:           true,
			},
			"api_token": schema.StringAttribute{
				MarkdownDescription: "ArcOrch api key",
				Optional:            true,
				Sensitive:           true,
			},
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "ArcOrch url",
				Required:            true,
			},
			"port": schema.StringAttribute{
				MarkdownDescription: "port",
				Required:            true,
			},
		},
	}
}

// Configure satisfies the provider.Provider interface for ExampleCloudProvider.
func (p *ArcOrchProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ArcOrchProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	loginCred := utils.LoginCred{
		BaseUrl: fmt.Sprintf(`%s:%s/api/v1/`, data.Endpoint.ValueString(), data.Port.ValueString()),
	}

	accessToken := ""
	authHeader := ""
	tenant := ""

	if data.Endpoint.IsNull() || data.Endpoint.IsUnknown() || data.Endpoint.ValueString() == "" {
		resp.Diagnostics.AddError("Missing Endpoint Configuration", "Missing Endpoint Configuration")
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() && data.Username.ValueString() != "" && !data.Password.IsNull() {
		username := data.Username.ValueString()
		tokens := strings.Split(username, "@")
		if len(tokens) == 1 {
			username = username + "@arcorch.com"
			tenant = "arcorch.com"
		} else if len(tokens) == 2 {
			tenant = tokens[1]
		} else {
			resp.Diagnostics.AddError("invalid username", "invalid username")
			return
		}
		authHeader = "Authorization"

		loginUrl := fmt.Sprintf(`%s:%s/api/v1/login?tenant=%s`, data.Endpoint.ValueString(), data.Port.ValueString(), tenant)
		user := map[string]string{
			"username": username,
			"tenant":   tenant,
			"password": data.Password.ValueString(),
		}

		loginResp, err := utils.PostRequest(loginUrl, user, "", "")
		if err != nil {
			resp.Diagnostics.AddError("Failed to login", err.Error())
			return
		}
		if loginResp.StatusCode != 200 {
			body, _ := ioutil.ReadAll(loginResp.Body)
			sb := string(body)
			resp.Diagnostics.AddError("Failed to login", sb)
			return
		}
		var result map[string]string
		if err := json.NewDecoder(loginResp.Body).Decode(&result); err != nil {
			resp.Diagnostics.AddError("Failed to login", err.Error())
			return
		}

		accessToken = result["access_token"]
	} else if !data.ApiKey.IsNull() && !data.ApiKey.IsUnknown() && data.ApiKey.ValueString() != "" {
		tokens := strings.Split(data.ApiKey.ValueString(), ".")
		if len(tokens) == 2 {
			strByte, err := b64.StdEncoding.DecodeString(tokens[0])
			if err != nil {
				resp.Diagnostics.AddError("invalid api key", err.Error())
				return
			}
			tenant = string(strByte)
		} else {
			resp.Diagnostics.AddError("invalid api key", "invalid api key")
			return
		}

		authHeader = "ArcOrchApiKey"
		accessToken = data.ApiKey.ValueString()
	} else {
		resp.Diagnostics.AddError("Missing ApiKey or Username/password Configuration", "Missing ApiKey or Username/password Configuration")
		return
	}

	loginCred.AccessToken = accessToken
	loginCred.AuthorizationHeader = authHeader
	loginCred.Tenant = tenant

	resp.DataSourceData = loginCred
	resp.ResourceData = loginCred
}

// DataSources satisfies the provider.Provider interface for ExampleCloudProvider.
func (p *ArcOrchProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datacredential.NewAwsCredDataSource,
		datacredential.NewAzureCredDataSource,
		datacredential.NewGcpCredDataSource,
		datacredential.NewOciCredDataSource,
		datacredential.NewKvmCredDataSource,
		datadeployment.NewAwsDeploymentDataSource,
		datadeployment.NewAzureDeploymentDataSource,
		datadeployment.NewGcpDeploymentDataSource,
		datadeployment.NewOciDeploymentDataSource,
		datadeployment.NewKvmDeploymentDataSource,
	}
}

// Resources satisfies the provider.Provider interface for ExampleCloudProvider.
func (p *ArcOrchProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		arc_resource.NewUserResource,
		credentialresource.NewAwsCredentialResource,
		credentialresource.NewAzureCredentialResource,
		credentialresource.NewGcpCredentialResource,
		credentialresource.NewOciCredentialResource,
		credentialresource.NewKvmCredentialResource,
		deploymentresource.NewAwsDeploymentResource,
		deploymentresource.NewAzureDeploymentResource,
		deploymentresource.NewGcpDeploymentResource,
		deploymentresource.NewOciDeploymentResource,
		deploymentresource.NewKvmDeploymentResource,
	}
}
