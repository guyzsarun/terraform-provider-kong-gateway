// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"net/http"
)

var _ provider.Provider = &KongGatewayProvider{}

type KongGatewayProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KongGatewayProviderModel describes the provider data model.
type KongGatewayProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
}

func (p *KongGatewayProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "kong-gateway"
	resp.Version = p.version
}

func (p *KongGatewayProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Kong Gateway Admin API: OpenAPI 3.0 spec for Kong Gateway's Admin API.` + "\n" +
			`` + "\n" +
			`You can lean more about Kong Gateway at [docs.konghq.com](https://docs.konghq.com)` + "\n" +
			`.Give Kong a star at [Kong/kong](https://github.com/kong/kong) repository.`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to {protocol}://{hostname}:{port}{path})",
				Optional:            true,
				Required:            false,
			},
		},
	}
}

func (p *KongGatewayProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KongGatewayProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "{protocol}://{hostname}:{port}{path}"
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *KongGatewayProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewACLResource,
		NewBasicAuthResource,
		NewCACertificateResource,
		NewCertificateResource,
		NewConsumerResource,
		NewConsumerGroupResource,
		NewGatewayConsumerGroupMemberResource,
		NewHMACAuthResource,
		NewJwtResource,
		NewKeyResource,
		NewKeyAuthResource,
		NewKeySetResource,
		NewPluginACLResource,
		NewPluginAcmeResource,
		NewPluginAiPromptDecoratorResource,
		NewPluginAiPromptGuardResource,
		NewPluginAiPromptTemplateResource,
		NewPluginAiProxyResource,
		NewPluginAiRequestTransformerResource,
		NewPluginAiResponseTransformerResource,
		NewPluginAwsLambdaResource,
		NewPluginAzureFunctionsResource,
		NewPluginBasicAuthResource,
		NewPluginBotDetectionResource,
		NewPluginCanaryResource,
		NewPluginCorrelationIDResource,
		NewPluginCorsResource,
		NewPluginDatadogResource,
		NewPluginDegraphqlResource,
		NewPluginExitTransformerResource,
		NewPluginFileLogResource,
		NewPluginForwardProxyResource,
		NewPluginGraphqlProxyCacheAdvancedResource,
		NewPluginGraphqlRateLimitingAdvancedResource,
		NewPluginGrpcGatewayResource,
		NewPluginGrpcWebResource,
		NewPluginHmacAuthResource,
		NewPluginHTTPLogResource,
		NewPluginIPRestrictionResource,
		NewPluginJqResource,
		NewPluginJweDecryptResource,
		NewPluginJwtResource,
		NewPluginJwtSignerResource,
		NewPluginKafkaLogResource,
		NewPluginKafkaUpstreamResource,
		NewPluginKeyAuthResource,
		NewPluginKeyAuthEncResource,
		NewPluginKonnectApplicationAuthResource,
		NewPluginLdapAuthResource,
		NewPluginLdapAuthAdvancedResource,
		NewPluginLogglyResource,
		NewPluginMockingResource,
		NewPluginMtlsAuthResource,
		NewPluginOasValidationResource,
		NewPluginOauth2Resource,
		NewPluginOauth2IntrospectionResource,
		NewPluginOpaResource,
		NewPluginOpenidConnectResource,
		NewPluginOpentelemetryResource,
		NewPluginPostFunctionResource,
		NewPluginPreFunctionResource,
		NewPluginPrometheusResource,
		NewPluginProxyCacheResource,
		NewPluginProxyCacheAdvancedResource,
		NewPluginRateLimitingResource,
		NewPluginRateLimitingAdvancedResource,
		NewPluginRequestSizeLimitingResource,
		NewPluginRequestTerminationResource,
		NewPluginRequestTransformerResource,
		NewPluginRequestTransformerAdvancedResource,
		NewPluginRequestValidatorResource,
		NewPluginResponseRatelimitingResource,
		NewPluginResponseTransformerResource,
		NewPluginResponseTransformerAdvancedResource,
		NewPluginRouteByHeaderResource,
		NewPluginRouteTransformerAdvancedResource,
		NewPluginSamlResource,
		NewPluginSessionResource,
		NewPluginStatsdResource,
		NewPluginStatsdAdvancedResource,
		NewPluginSyslogResource,
		NewPluginTCPLogResource,
		NewPluginTLSHandshakeModifierResource,
		NewPluginTLSMetadataHeadersResource,
		NewPluginUDPLogResource,
		NewPluginUpstreamTimeoutResource,
		NewPluginVaultAuthResource,
		NewPluginWebsocketSizeLimitResource,
		NewPluginWebsocketValidatorResource,
		NewPluginXMLThreatProtectionResource,
		NewPluginZipkinResource,
		NewRouteResource,
		NewServiceResource,
		NewSniResource,
		NewUpstreamResource,
		NewVaultResource,
	}
}

func (p *KongGatewayProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewACLDataSource,
		NewBasicAuthDataSource,
		NewCACertificateDataSource,
		NewCertificateDataSource,
		NewConsumerDataSource,
		NewConsumerGroupDataSource,
		NewHMACAuthDataSource,
		NewJwtDataSource,
		NewKeyDataSource,
		NewKeyAuthDataSource,
		NewKeySetDataSource,
		NewPluginACLDataSource,
		NewPluginAcmeDataSource,
		NewPluginAiPromptDecoratorDataSource,
		NewPluginAiPromptGuardDataSource,
		NewPluginAiPromptTemplateDataSource,
		NewPluginAiProxyDataSource,
		NewPluginAiRequestTransformerDataSource,
		NewPluginAiResponseTransformerDataSource,
		NewPluginAwsLambdaDataSource,
		NewPluginAzureFunctionsDataSource,
		NewPluginBasicAuthDataSource,
		NewPluginBotDetectionDataSource,
		NewPluginCanaryDataSource,
		NewPluginCorrelationIDDataSource,
		NewPluginCorsDataSource,
		NewPluginDatadogDataSource,
		NewPluginDegraphqlDataSource,
		NewPluginExitTransformerDataSource,
		NewPluginFileLogDataSource,
		NewPluginForwardProxyDataSource,
		NewPluginGraphqlProxyCacheAdvancedDataSource,
		NewPluginGraphqlRateLimitingAdvancedDataSource,
		NewPluginGrpcGatewayDataSource,
		NewPluginGrpcWebDataSource,
		NewPluginHmacAuthDataSource,
		NewPluginHTTPLogDataSource,
		NewPluginIPRestrictionDataSource,
		NewPluginJqDataSource,
		NewPluginJweDecryptDataSource,
		NewPluginJwtDataSource,
		NewPluginJwtSignerDataSource,
		NewPluginKafkaLogDataSource,
		NewPluginKafkaUpstreamDataSource,
		NewPluginKeyAuthDataSource,
		NewPluginKeyAuthEncDataSource,
		NewPluginKonnectApplicationAuthDataSource,
		NewPluginLdapAuthDataSource,
		NewPluginLdapAuthAdvancedDataSource,
		NewPluginLogglyDataSource,
		NewPluginMockingDataSource,
		NewPluginMtlsAuthDataSource,
		NewPluginOasValidationDataSource,
		NewPluginOauth2DataSource,
		NewPluginOauth2IntrospectionDataSource,
		NewPluginOpaDataSource,
		NewPluginOpenidConnectDataSource,
		NewPluginOpentelemetryDataSource,
		NewPluginPostFunctionDataSource,
		NewPluginPreFunctionDataSource,
		NewPluginPrometheusDataSource,
		NewPluginProxyCacheDataSource,
		NewPluginProxyCacheAdvancedDataSource,
		NewPluginRateLimitingDataSource,
		NewPluginRateLimitingAdvancedDataSource,
		NewPluginRequestSizeLimitingDataSource,
		NewPluginRequestTerminationDataSource,
		NewPluginRequestTransformerDataSource,
		NewPluginRequestTransformerAdvancedDataSource,
		NewPluginRequestValidatorDataSource,
		NewPluginResponseRatelimitingDataSource,
		NewPluginResponseTransformerDataSource,
		NewPluginResponseTransformerAdvancedDataSource,
		NewPluginRouteByHeaderDataSource,
		NewPluginRouteTransformerAdvancedDataSource,
		NewPluginSamlDataSource,
		NewPluginSessionDataSource,
		NewPluginStatsdDataSource,
		NewPluginStatsdAdvancedDataSource,
		NewPluginSyslogDataSource,
		NewPluginTCPLogDataSource,
		NewPluginTLSHandshakeModifierDataSource,
		NewPluginTLSMetadataHeadersDataSource,
		NewPluginUDPLogDataSource,
		NewPluginUpstreamTimeoutDataSource,
		NewPluginVaultAuthDataSource,
		NewPluginWebsocketSizeLimitDataSource,
		NewPluginWebsocketValidatorDataSource,
		NewPluginXMLThreatProtectionDataSource,
		NewPluginZipkinDataSource,
		NewRouteDataSource,
		NewServiceDataSource,
		NewSniDataSource,
		NewUpstreamDataSource,
		NewVaultDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KongGatewayProvider{
			version: version,
		}
	}
}
