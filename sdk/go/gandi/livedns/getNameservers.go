// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package livedns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNameservers(ctx *pulumi.Context, args *GetNameserversArgs, opts ...pulumi.InvokeOption) (*GetNameserversResult, error) {
	var rv GetNameserversResult
	err := ctx.Invoke("gandi:livedns/getNameservers:getNameservers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNameservers.
type GetNameserversArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getNameservers.
type GetNameserversResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	Name        string   `pulumi:"name"`
	Nameservers []string `pulumi:"nameservers"`
}

func GetNameserversOutput(ctx *pulumi.Context, args GetNameserversOutputArgs, opts ...pulumi.InvokeOption) GetNameserversResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNameserversResult, error) {
			args := v.(GetNameserversArgs)
			r, err := GetNameservers(ctx, &args, opts...)
			return *r, err
		}).(GetNameserversResultOutput)
}

// A collection of arguments for invoking getNameservers.
type GetNameserversOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetNameserversOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNameserversArgs)(nil)).Elem()
}

// A collection of values returned by getNameservers.
type GetNameserversResultOutput struct{ *pulumi.OutputState }

func (GetNameserversResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNameserversResult)(nil)).Elem()
}

func (o GetNameserversResultOutput) ToGetNameserversResultOutput() GetNameserversResultOutput {
	return o
}

func (o GetNameserversResultOutput) ToGetNameserversResultOutputWithContext(ctx context.Context) GetNameserversResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNameserversResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNameserversResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNameserversResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNameserversResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNameserversResultOutput) Nameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNameserversResult) []string { return v.Nameservers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNameserversResultOutput{})
}