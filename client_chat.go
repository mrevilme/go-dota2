package dota2

import (
	gcm "github.com/paralin/go-dota2/protocol"
	gcmcc "github.com/paralin/go-dota2/protocol"

	"context"
	"fmt"
)

// SendChannelMessage attempts to send a message in a channel, text-only.
// Use SendChatMessage for more fine-grained control.
func (d *Dota2) SendChannelMessage(channelID uint64, message string) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCChatMessage), &gcmcc.CMsgDOTAChatMessage{
		ChannelId: &channelID,
		Text:      &message,
	})
}

func (d *Dota2) JoinLobbyChat(ctx context.Context, lobby *gcm.CSODOTALobby) (*gcm.CMsgDOTAJoinChatChannelResponse, error) {
	return d.JoinChatChannel(ctx, fmt.Sprintf("Lobby_%d", *lobby.LobbyId), gcm.DOTAChatChannelTypeT_DOTAChannelType_Lobby, true)
}
