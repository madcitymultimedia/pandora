using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.WebPubSub.v2023_02_01.WebPubSub;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ScaleTypeConstant
{
    [Description("Automatic")]
    Automatic,

    [Description("Manual")]
    Manual,

    [Description("None")]
    None,
}
