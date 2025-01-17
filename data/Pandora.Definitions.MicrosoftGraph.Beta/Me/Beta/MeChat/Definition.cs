// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

using Pandora.Definitions.Interfaces;
using Pandora.Definitions.MicrosoftGraph.Beta.CommonTypes;
using System;

namespace Pandora.Definitions.MicrosoftGraph.Beta.Me.Beta.MeChat;

internal class Definition : ResourceDefinition
{
    public string Name => "MeChat";

    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateMeChatByIdHideForUserOperation(),
        new CreateMeChatByIdMarkChatReadForUserOperation(),
        new CreateMeChatByIdMarkChatUnreadForUserOperation(),
        new CreateMeChatByIdSendActivityNotificationOperation(),
        new CreateMeChatByIdUnhideForUserOperation(),
        new CreateMeChatOperation(),
        new DeleteMeChatByIdOperation(),
        new GetMeChatByIdOperation(),
        new GetMeChatCountOperation(),
        new ListMeChatsOperation(),
        new RemoveMeChatByIdAllAccessForUserOperation(),
        new UpdateMeChatByIdOperation()
    };

    public IEnumerable<System.Type> Constants => new List<System.Type>
    {

    };

    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(CreateMeChatByIdHideForUserRequestModel),
        typeof(CreateMeChatByIdMarkChatReadForUserRequestModel),
        typeof(CreateMeChatByIdMarkChatUnreadForUserRequestModel),
        typeof(CreateMeChatByIdSendActivityNotificationRequestModel),
        typeof(CreateMeChatByIdUnhideForUserRequestModel),
        typeof(RemoveMeChatByIdAllAccessForUserRequestModel)
    };
}
