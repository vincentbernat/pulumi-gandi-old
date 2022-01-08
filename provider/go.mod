module github.com/vincentbernat/pulumi-gandi/provider

go 1.16

replace (
	github.com/go-gandi/go-gandi v0.0.0-20211230165416-1a297fd86f69 => github.com/vincentbernat/go-gandi v0.0.0-20220108220653-c49b75534f6f
	github.com/go-gandi/terraform-provider-gandi v1.1.2-0.20211230172315-5378a78616de => github.com/vincentbernat/terraform-provider-gandi v1.1.2-0.20220108221411-d073d00f7ef0
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
)

require (
	github.com/go-gandi/terraform-provider-gandi v1.1.2-0.20211230172315-5378a78616de
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.13.0
	github.com/pulumi/pulumi/sdk/v3 v3.19.0
)
