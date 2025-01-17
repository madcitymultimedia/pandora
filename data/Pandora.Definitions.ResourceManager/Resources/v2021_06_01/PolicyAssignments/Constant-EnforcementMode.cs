using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Resources.v2021_06_01.PolicyAssignments;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum EnforcementModeConstant
{
    [Description("Default")]
    Default,

    [Description("DoNotEnforce")]
    DoNotEnforce,
}
