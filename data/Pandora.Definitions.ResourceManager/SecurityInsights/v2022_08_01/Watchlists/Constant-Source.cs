using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.SecurityInsights.v2022_08_01.Watchlists;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum SourceConstant
{
    [Description("Local file")]
    LocalFile,

    [Description("Remote storage")]
    RemoteStorage,
}
