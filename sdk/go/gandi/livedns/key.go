// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package livedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Key struct {
	pulumi.CustomResourceState

	// DNSSEC algorithm type
	Algorithm pulumi.IntOutput `pulumi:"algorithm"`
	// DNSSEC algorithm name
	AlgorithmName pulumi.StringOutput `pulumi:"algorithmName"`
	// Is the key deleted?
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// Domain name
	Domain pulumi.StringOutput `pulumi:"domain"`
	// DS record as RFC1035 line
	Ds pulumi.StringOutput `pulumi:"ds"`
	// DNSSEC key flags
	Flags pulumi.IntOutput `pulumi:"flags"`
	// Public key
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// Current status of the key
	Status pulumi.StringOutput `pulumi:"status"`
	// Tag
	Tag pulumi.IntOutput `pulumi:"tag"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource Key
	err := ctx.RegisterResource("gandi:livedns/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("gandi:livedns/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// DNSSEC algorithm type
	Algorithm *int `pulumi:"algorithm"`
	// DNSSEC algorithm name
	AlgorithmName *string `pulumi:"algorithmName"`
	// Is the key deleted?
	Deleted *bool `pulumi:"deleted"`
	// Domain name
	Domain *string `pulumi:"domain"`
	// DS record as RFC1035 line
	Ds *string `pulumi:"ds"`
	// DNSSEC key flags
	Flags *int `pulumi:"flags"`
	// Public key
	PublicKey *string `pulumi:"publicKey"`
	// Current status of the key
	Status *string `pulumi:"status"`
	// Tag
	Tag *int `pulumi:"tag"`
}

type KeyState struct {
	// DNSSEC algorithm type
	Algorithm pulumi.IntPtrInput
	// DNSSEC algorithm name
	AlgorithmName pulumi.StringPtrInput
	// Is the key deleted?
	Deleted pulumi.BoolPtrInput
	// Domain name
	Domain pulumi.StringPtrInput
	// DS record as RFC1035 line
	Ds pulumi.StringPtrInput
	// DNSSEC key flags
	Flags pulumi.IntPtrInput
	// Public key
	PublicKey pulumi.StringPtrInput
	// Current status of the key
	Status pulumi.StringPtrInput
	// Tag
	Tag pulumi.IntPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Domain name
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Domain name
	Domain pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//          KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//          KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Key {
		return vs[0].([]*Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Key {
		return vs[0].(map[string]*Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyArrayInput)(nil)).Elem(), KeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyMapInput)(nil)).Elem(), KeyMap{})
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
