using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Compute.v2023_04_02.DiskAccesses;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum PrivateEndpointServiceConnectionStatusConstant
{
    [Description("Approved")]
    Approved,

    [Description("Pending")]
    Pending,

    [Description("Rejected")]
    Rejected,
}
