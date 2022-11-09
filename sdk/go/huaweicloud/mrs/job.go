// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mrs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a job resource within HuaweiCloud MRS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Mrs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			jobName := cfg.RequireObject("jobName")
//			programPath := cfg.RequireObject("programPath")
//			accessKey := cfg.RequireObject("accessKey")
//			secretKey := cfg.RequireObject("secretKey")
//			_, err := Mrs.NewJob(ctx, "test", &Mrs.JobArgs{
//				ClusterId:   pulumi.Any(clusterId),
//				Type:        pulumi.String("SparkSubmit"),
//				ProgramPath: pulumi.Any(programPath),
//				Parameters:  pulumi.String(fmt.Sprintf("%v %v 1 obs://obs-demo-analysis/input obs://obs-demo-analysis/output", accessKey, secretKey)),
//				ProgramParameters: pulumi.StringMap{
//					"--class": pulumi.String("com.huawei.bigdata.spark.examples.DriverBehavior"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// MapReduce jobs can be imported using their `id` and the IDs of the MapReduce cluster to which the job belongs, separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Mrs/job:Job test <cluster_id>/<id>
//
// ```
type Job struct {
	pulumi.CustomResourceState

	// Specifies an ID of the MapReduce cluster to which the job belongs to.
	// Changing this will create a new MapReduce job resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The completion time of the MapReduce job.
	FinishTime pulumi.StringOutput `pulumi:"finishTime"`
	// Specifies the name of the MapReduce job. The name can contain 1 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MapReduce job resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the parameters for the MapReduce job. Add an at sign (@) before
	// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
	// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
	// MapReduce job resource.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// Specifies the the key/value pairs of the program parameters, such as
	// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
	// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
	// **HiveScript**. Refer to the documents for each type of support key-values.
	// Changing this will create a new MapReduce job resource.
	ProgramParameters pulumi.StringMapOutput `pulumi:"programParameters"`
	// Specifies the .jar package path or .py file path for program execution.
	// The parameter must meet the following requirements:
	// + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
	// + The address cannot be empty or full of spaces.
	// + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
	//   with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
	ProgramPath pulumi.StringPtrOutput `pulumi:"programPath"`
	// Specifies the region in which to create the MapReduce job resource. If
	// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the key/value pairs used to modify service configuration.
	// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
	// Changing this will create a new MapReduce job resource.
	ServiceParameters pulumi.StringMapOutput `pulumi:"serviceParameters"`
	// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
	// or **SparkSql**. Changing this will create a new MapReduce job resource.
	Sql pulumi.StringPtrOutput `pulumi:"sql"`
	// The creation time of the MapReduce job.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Status of the MapReduce job.
	Status pulumi.StringOutput `pulumi:"status"`
	// The submission time of the MapReduce job.
	SubmitTime pulumi.StringOutput `pulumi:"submitTime"`
	Type       pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("huaweicloud:Mrs/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("huaweicloud:Mrs/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Specifies an ID of the MapReduce cluster to which the job belongs to.
	// Changing this will create a new MapReduce job resource.
	ClusterId *string `pulumi:"clusterId"`
	// The completion time of the MapReduce job.
	FinishTime *string `pulumi:"finishTime"`
	// Specifies the name of the MapReduce job. The name can contain 1 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MapReduce job resource.
	Name *string `pulumi:"name"`
	// Specifies the parameters for the MapReduce job. Add an at sign (@) before
	// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
	// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
	// MapReduce job resource.
	Parameters *string `pulumi:"parameters"`
	// Specifies the the key/value pairs of the program parameters, such as
	// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
	// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
	// **HiveScript**. Refer to the documents for each type of support key-values.
	// Changing this will create a new MapReduce job resource.
	ProgramParameters map[string]string `pulumi:"programParameters"`
	// Specifies the .jar package path or .py file path for program execution.
	// The parameter must meet the following requirements:
	// + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
	// + The address cannot be empty or full of spaces.
	// + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
	//   with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
	ProgramPath *string `pulumi:"programPath"`
	// Specifies the region in which to create the MapReduce job resource. If
	// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs used to modify service configuration.
	// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
	// Changing this will create a new MapReduce job resource.
	ServiceParameters map[string]string `pulumi:"serviceParameters"`
	// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
	// or **SparkSql**. Changing this will create a new MapReduce job resource.
	Sql *string `pulumi:"sql"`
	// The creation time of the MapReduce job.
	StartTime *string `pulumi:"startTime"`
	// Status of the MapReduce job.
	Status *string `pulumi:"status"`
	// The submission time of the MapReduce job.
	SubmitTime *string `pulumi:"submitTime"`
	Type       *string `pulumi:"type"`
}

type JobState struct {
	// Specifies an ID of the MapReduce cluster to which the job belongs to.
	// Changing this will create a new MapReduce job resource.
	ClusterId pulumi.StringPtrInput
	// The completion time of the MapReduce job.
	FinishTime pulumi.StringPtrInput
	// Specifies the name of the MapReduce job. The name can contain 1 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MapReduce job resource.
	Name pulumi.StringPtrInput
	// Specifies the parameters for the MapReduce job. Add an at sign (@) before
	// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
	// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
	// MapReduce job resource.
	Parameters pulumi.StringPtrInput
	// Specifies the the key/value pairs of the program parameters, such as
	// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
	// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
	// **HiveScript**. Refer to the documents for each type of support key-values.
	// Changing this will create a new MapReduce job resource.
	ProgramParameters pulumi.StringMapInput
	// Specifies the .jar package path or .py file path for program execution.
	// The parameter must meet the following requirements:
	// + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
	// + The address cannot be empty or full of spaces.
	// + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
	//   with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
	ProgramPath pulumi.StringPtrInput
	// Specifies the region in which to create the MapReduce job resource. If
	// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs used to modify service configuration.
	// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
	// Changing this will create a new MapReduce job resource.
	ServiceParameters pulumi.StringMapInput
	// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
	// or **SparkSql**. Changing this will create a new MapReduce job resource.
	Sql pulumi.StringPtrInput
	// The creation time of the MapReduce job.
	StartTime pulumi.StringPtrInput
	// Status of the MapReduce job.
	Status pulumi.StringPtrInput
	// The submission time of the MapReduce job.
	SubmitTime pulumi.StringPtrInput
	Type       pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Specifies an ID of the MapReduce cluster to which the job belongs to.
	// Changing this will create a new MapReduce job resource.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the name of the MapReduce job. The name can contain 1 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MapReduce job resource.
	Name *string `pulumi:"name"`
	// Specifies the parameters for the MapReduce job. Add an at sign (@) before
	// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
	// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
	// MapReduce job resource.
	Parameters *string `pulumi:"parameters"`
	// Specifies the the key/value pairs of the program parameters, such as
	// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
	// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
	// **HiveScript**. Refer to the documents for each type of support key-values.
	// Changing this will create a new MapReduce job resource.
	ProgramParameters map[string]string `pulumi:"programParameters"`
	// Specifies the .jar package path or .py file path for program execution.
	// The parameter must meet the following requirements:
	// + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
	// + The address cannot be empty or full of spaces.
	// + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
	//   with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
	ProgramPath *string `pulumi:"programPath"`
	// Specifies the region in which to create the MapReduce job resource. If
	// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs used to modify service configuration.
	// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
	// Changing this will create a new MapReduce job resource.
	ServiceParameters map[string]string `pulumi:"serviceParameters"`
	// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
	// or **SparkSql**. Changing this will create a new MapReduce job resource.
	Sql  *string `pulumi:"sql"`
	Type string  `pulumi:"type"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Specifies an ID of the MapReduce cluster to which the job belongs to.
	// Changing this will create a new MapReduce job resource.
	ClusterId pulumi.StringInput
	// Specifies the name of the MapReduce job. The name can contain 1 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MapReduce job resource.
	Name pulumi.StringPtrInput
	// Specifies the parameters for the MapReduce job. Add an at sign (@) before
	// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
	// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
	// MapReduce job resource.
	Parameters pulumi.StringPtrInput
	// Specifies the the key/value pairs of the program parameters, such as
	// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
	// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
	// **HiveScript**. Refer to the documents for each type of support key-values.
	// Changing this will create a new MapReduce job resource.
	ProgramParameters pulumi.StringMapInput
	// Specifies the .jar package path or .py file path for program execution.
	// The parameter must meet the following requirements:
	// + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
	// + The address cannot be empty or full of spaces.
	// + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
	//   with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
	ProgramPath pulumi.StringPtrInput
	// Specifies the region in which to create the MapReduce job resource. If
	// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs used to modify service configuration.
	// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
	// Changing this will create a new MapReduce job resource.
	ServiceParameters pulumi.StringMapInput
	// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
	// or **SparkSql**. Changing this will create a new MapReduce job resource.
	Sql  pulumi.StringPtrInput
	Type pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//	JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//	JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// Specifies an ID of the MapReduce cluster to which the job belongs to.
// Changing this will create a new MapReduce job resource.
func (o JobOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The completion time of the MapReduce job.
func (o JobOutput) FinishTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.FinishTime }).(pulumi.StringOutput)
}

// Specifies the name of the MapReduce job. The name can contain 1 to 64
// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
// MapReduce job resource.
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the parameters for the MapReduce job. Add an at sign (@) before
// each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
// This parameter can be set when `type` is **Flink**, **MapReduce** or **SparkSubmit**. Changing this will create a new
// MapReduce job resource.
func (o JobOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Parameters }).(pulumi.StringPtrOutput)
}

// Specifies the the key/value pairs of the program parameters, such as
// thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
// can be set when `type` is **Flink**, **SparkSubmit**, **SparkSql**, **SparkScript**, **HiveSql** or
// **HiveScript**. Refer to the documents for each type of support key-values.
// Changing this will create a new MapReduce job resource.
func (o JobOutput) ProgramParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.ProgramParameters }).(pulumi.StringMapOutput)
}

// Specifies the .jar package path or .py file path for program execution.
// The parameter must meet the following requirements:
//   - Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
//   - The address cannot be empty or full of spaces.
//   - The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
//     with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.
func (o JobOutput) ProgramPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.ProgramPath }).(pulumi.StringPtrOutput)
}

// Specifies the region in which to create the MapReduce job resource. If
// omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.
func (o JobOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the key/value pairs used to modify service configuration.
// Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
// Changing this will create a new MapReduce job resource.
func (o JobOutput) ServiceParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.ServiceParameters }).(pulumi.StringMapOutput)
}

// Specifies the SQL command or file path. Only required if `type` is **HiveSql**
// or **SparkSql**. Changing this will create a new MapReduce job resource.
func (o JobOutput) Sql() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Sql }).(pulumi.StringPtrOutput)
}

// The creation time of the MapReduce job.
func (o JobOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Status of the MapReduce job.
func (o JobOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The submission time of the MapReduce job.
func (o JobOutput) SubmitTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.SubmitTime }).(pulumi.StringOutput)
}

func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
