package dota2

import (
	"context"
	"github.com/Philipp15b/go-steam/v3/steamid"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol"
)

// AbandonLobby abandons a lobby.
// Request ID: k_EMsgGCAbandonCurrentGame
// Request type: CMsgAbandonCurrentGame
func (d *Dota2) AbandonLobby() {
	req := &protocol.CMsgAbandonCurrentGame{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame), req)
}

// AckPartyReadyCheck checks for/from a ack party ready.
// Request ID: k_EMsgPartyReadyCheckAcknowledge
// Request type: CMsgPartyReadyCheckAcknowledge
func (d *Dota2) AckPartyReadyCheck(
	readyStatus protocol.EReadyCheckStatus,
) {
	req := &protocol.CMsgPartyReadyCheckAcknowledge{
		ReadyStatus: &readyStatus,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge), req)
}

// ApplyGauntletTicket applys a gauntlet ticket.
// Request ID: k_EMsgClientToGCApplyGauntletTicket
// Request type: CMsgClientToGCApplyGauntletTicket
func (d *Dota2) ApplyGauntletTicket() {
	req := &protocol.CMsgClientToGCApplyGauntletTicket{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCApplyGauntletTicket), req)
}

// ApplyGemCombiner applys a gem combiner.
// Request ID: k_EMsgClientToGCApplyGemCombiner
// Request type: CMsgClientToGCApplyGemCombiner
func (d *Dota2) ApplyGemCombiner(
	itemID1 uint64,
	itemID2 uint64,
) {
	req := &protocol.CMsgClientToGCApplyGemCombiner{
		ItemId_1: &itemID1,
		ItemId_2: &itemID2,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCApplyGemCombiner), req)
}

// ApplyTeamToLobby applys a team to lobby.
// Request ID: k_EMsgGCApplyTeamToPracticeLobby
// Request type: CMsgApplyTeamToPracticeLobby
func (d *Dota2) ApplyTeamToLobby(
	teamID uint32,
) {
	req := &protocol.CMsgApplyTeamToPracticeLobby{
		TeamId: &teamID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCApplyTeamToPracticeLobby), req)
}

// AutographReward autographs a reward.
// Request ID: k_EMsgGameAutographReward
// Response ID: k_EMsgGameAutographRewardResponse
// Request type: CMsgDOTAGameAutographReward
// Response type: CMsgDOTAGameAutographRewardResponse
func (d *Dota2) AutographReward(
	ctx context.Context,
	badgeID string,
) (*protocol.CMsgDOTAGameAutographRewardResponse, error) {
	req := &protocol.CMsgDOTAGameAutographReward{
		BadgeId: &badgeID,
	}
	resp := &protocol.CMsgDOTAGameAutographRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGameAutographReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGameAutographRewardResponse),
		resp,
	)
}

// CancelInviteToGuild cancels a invite to guild.
// Request ID: k_EMsgClientToGCCancelInviteToGuild
// Response ID: k_EMsgClientToGCCancelInviteToGuildResponse
// Request type: CMsgClientToGCCancelInviteToGuild
// Response type: CMsgClientToGCCancelInviteToGuildResponse
func (d *Dota2) CancelInviteToGuild(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCCancelInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCCancelInviteToGuild{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCCancelInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelInviteToGuildResponse),
		resp,
	)
}

// CancelPartyInvites cancels party invites.
// Request ID: k_EMsgClientToGCCancelPartyInvites
// Request type: CMsgDOTACancelGroupInvites
func (d *Dota2) CancelPartyInvites(
	invitedSteamids []uint64,
	invitedGroupids []uint64,
) {
	req := &protocol.CMsgDOTACancelGroupInvites{
		InvitedSteamids: invitedSteamids,
		InvitedGroupids: invitedGroupids,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites), req)
}

// CancelWatchGame cancels a watch game.
// Request ID: k_EMsgGCCancelWatchGame
// Request type: CMsgCancelWatchGame
func (d *Dota2) CancelWatchGame() {
	req := &protocol.CMsgCancelWatchGame{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCCancelWatchGame), req)
}

// CloseLobbyBroadcastChannel closes a lobby broadcast channel.
// Request ID: k_EMsgGCPracticeLobbyCloseBroadcastChannel
// Request type: CMsgPracticeLobbyCloseBroadcastChannel
func (d *Dota2) CloseLobbyBroadcastChannel(
	channel uint32,
) {
	req := &protocol.CMsgPracticeLobbyCloseBroadcastChannel{
		Channel: &channel,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyCloseBroadcastChannel), req)
}

// CreateBotGame creates a bot game.
// Request ID: k_EMsgGCBotGameCreate
// Request type: CMsgBotGameCreate
func (d *Dota2) CreateBotGame(
	req *protocol.CMsgBotGameCreate,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCBotGameCreate), req)
}

// CreateGameCustom creates a game custom.
// Request ID: k_EMsgGCCustomGameCreate
// Request type: CMsgCustomGameCreate
func (d *Dota2) CreateGameCustom(
	req *protocol.CMsgCustomGameCreate,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCCustomGameCreate), req)
}

// CreateGuild creates a guild.
// Request ID: k_EMsgClientToGCCreateGuild
// Response ID: k_EMsgClientToGCCreateGuildResponse
// Request type: CMsgClientToGCCreateGuild
// Response type: CMsgClientToGCCreateGuildResponse
func (d *Dota2) CreateGuild(
	ctx context.Context,
	req *protocol.CMsgClientToGCCreateGuild,
) (*protocol.CMsgClientToGCCreateGuildResponse, error) {
	resp := &protocol.CMsgClientToGCCreateGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateGuildResponse),
		resp,
	)
}

// CreateHeroStatue creates a hero statue.
// Request ID: k_EMsgClientToGCCreateHeroStatue
// Request type: CMsgClientToGCCreateHeroStatue
func (d *Dota2) CreateHeroStatue(
	req *protocol.CMsgClientToGCCreateHeroStatue,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateHeroStatue), req)
}

// CreatePlayerCardPack creates a player card pack.
// Request ID: k_EMsgClientToGCCreatePlayerCardPack
// Response ID: k_EMsgClientToGCCreatePlayerCardPackResponse
// Request type: CMsgClientToGCCreatePlayerCardPack
// Response type: CMsgClientToGCCreatePlayerCardPackResponse
func (d *Dota2) CreatePlayerCardPack(
	ctx context.Context,
	cardDustItemID uint64,
	eventID uint32,
	premiumPack bool,
) (*protocol.CMsgClientToGCCreatePlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCCreatePlayerCardPack{
		CardDustItemId: &cardDustItemID,
		EventId:        &eventID,
		PremiumPack:    &premiumPack,
	}
	resp := &protocol.CMsgClientToGCCreatePlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse),
		resp,
	)
}

// CreateSpectatorLobby creates a spectator lobby.
// Request ID: k_EMsgClientToGCCreateSpectatorLobby
// Request type: CMsgCreateSpectatorLobby
func (d *Dota2) CreateSpectatorLobby(
	req *protocol.CMsgCreateSpectatorLobby,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobby), req)
}

// CreateTeam creates a team.
// Request ID: k_EMsgGCCreateTeam
// Response ID: k_EMsgGCCreateTeamResponse
// Request type: CMsgDOTACreateTeam
// Response type: CMsgDOTACreateTeamResponse
func (d *Dota2) CreateTeam(
	ctx context.Context,
	req *protocol.CMsgDOTACreateTeam,
) (*protocol.CMsgDOTACreateTeamResponse, error) {
	resp := &protocol.CMsgDOTACreateTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCreateTeam),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCreateTeamResponse),
		resp,
	)
}

// CreateTeamPlayerCardPack creates a team player card pack.
// Request ID: k_EMsgClientToGCCreateTeamPlayerCardPack
// Response ID: k_EMsgClientToGCCreateTeamPlayerCardPackResponse
// Request type: CMsgClientToGCCreateTeamPlayerCardPack
// Response type: CMsgClientToGCCreateTeamPlayerCardPackResponse
func (d *Dota2) CreateTeamPlayerCardPack(
	ctx context.Context,
	cardDustItemID uint64,
	eventID uint32,
	premiumPack bool,
	teamID uint32,
) (*protocol.CMsgClientToGCCreateTeamPlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCCreateTeamPlayerCardPack{
		CardDustItemId: &cardDustItemID,
		EventId:        &eventID,
		PremiumPack:    &premiumPack,
		TeamId:         &teamID,
	}
	resp := &protocol.CMsgClientToGCCreateTeamPlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateTeamPlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateTeamPlayerCardPackResponse),
		resp,
	)
}

// DemotePrivateChatMember demotes a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatDemote
// Request type: CMsgClientToGCPrivateChatDemote
func (d *Dota2) DemotePrivateChatMember(
	privateChatChannelName string,
	demoteAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatDemote{
		PrivateChatChannelName: &privateChatChannelName,
		DemoteAccountId:        &demoteAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote), req)
}

// DestroyLobby destroys a lobby.
// Request ID: k_EMsgDestroyLobbyRequest
// Response ID: k_EMsgDestroyLobbyResponse
// Request type: CMsgDOTADestroyLobbyRequest
// Response type: CMsgDOTADestroyLobbyResponse
func (d *Dota2) DestroyLobby(
	ctx context.Context,
) (*protocol.CMsgDOTADestroyLobbyResponse, error) {
	req := &protocol.CMsgDOTADestroyLobbyRequest{}
	resp := &protocol.CMsgDOTADestroyLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDestroyLobbyRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDestroyLobbyResponse),
		resp,
	)
}

// EditTeamDetails edits team details.
// Request ID: k_EMsgGCEditTeamDetails
// Response ID: k_EMsgGCEditTeamDetailsResponse
// Request type: CMsgDOTAEditTeamDetails
// Response type: CMsgDOTAEditTeamDetailsResponse
func (d *Dota2) EditTeamDetails(
	ctx context.Context,
	req *protocol.CMsgDOTAEditTeamDetails,
) (*protocol.CMsgDOTAEditTeamDetailsResponse, error) {
	resp := &protocol.CMsgDOTAEditTeamDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCEditTeamDetails),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCEditTeamDetailsResponse),
		resp,
	)
}

// FindTopSourceTVGames finds top source tv games.
// Request ID: k_EMsgClientToGCFindTopSourceTVGames
// Response ID: k_EMsgGCToClientFindTopSourceTVGamesResponse
// Request type: CMsgClientToGCFindTopSourceTVGames
// Response type: CMsgGCToClientFindTopSourceTVGamesResponse
func (d *Dota2) FindTopSourceTVGames(
	ctx context.Context,
	req *protocol.CMsgClientToGCFindTopSourceTVGames,
) (*protocol.CMsgGCToClientFindTopSourceTVGamesResponse, error) {
	resp := &protocol.CMsgGCToClientFindTopSourceTVGamesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse),
		resp,
	)
}

// FlipLobbyTeams flips lobby teams.
// Request ID: k_EMsgGCFlipLobbyTeams
// Request type: CMsgFlipLobbyTeams
func (d *Dota2) FlipLobbyTeams() {
	req := &protocol.CMsgFlipLobbyTeams{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCFlipLobbyTeams), req)
}

// GetAllHeroOrder gets all hero order.
// Request ID: k_EMsgClientToGCGetAllHeroOrder
// Response ID: k_EMsgClientToGCGetAllHeroOrderResponse
// Request type: CMsgClientToGCGetAllHeroOrder
// Response type: CMsgClientToGCGetAllHeroOrderResponse
func (d *Dota2) GetAllHeroOrder(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetAllHeroOrderResponse, error) {
	req := &protocol.CMsgClientToGCGetAllHeroOrder{}
	resp := &protocol.CMsgClientToGCGetAllHeroOrderResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrder),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrderResponse),
		resp,
	)
}

// GetAllHeroProgress gets all hero progress.
// Request ID: k_EMsgClientToGCGetAllHeroProgress
// Response ID: k_EMsgClientToGCGetAllHeroProgressResponse
// Request type: CMsgClientToGCGetAllHeroProgress
// Response type: CMsgClientToGCGetAllHeroProgressResponse
func (d *Dota2) GetAllHeroProgress(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetAllHeroProgressResponse, error) {
	req := &protocol.CMsgClientToGCGetAllHeroProgress{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetAllHeroProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgress),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgressResponse),
		resp,
	)
}

// GetAvailablePrivateCoachingSessions gets available private coaching sessions.
// Request ID: k_EMsgClientToGCGetAvailablePrivateCoachingSessions
// Response ID: k_EMsgClientToGCGetAvailablePrivateCoachingSessionsResponse
// Request type: CMsgClientToGCGetAvailablePrivateCoachingSessions
// Response type: CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse
func (d *Dota2) GetAvailablePrivateCoachingSessions(
	ctx context.Context,
	language uint32,
) (*protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse, error) {
	req := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessions{
		Language: &language,
	}
	resp := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsResponse),
		resp,
	)
}

// GetAvailablePrivateCoachingSessionsSummary gets a available private coaching sessions summary.
// Request ID: k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummary
// Response ID: k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse
// Request type: CMsgClientToGCGetAvailablePrivateCoachingSessionsSummary
// Response type: CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse
func (d *Dota2) GetAvailablePrivateCoachingSessionsSummary(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse, error) {
	req := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummary{}
	resp := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummary),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse),
		resp,
	)
}

// GetChatMemberCount gets a chat member count.
// Request ID: k_EMsgDOTAChatGetMemberCount
// Response ID: k_EMsgDOTAChatGetMemberCountResponse
// Request type: CMsgDOTAChatGetMemberCount
// Response type: CMsgDOTAChatGetMemberCountResponse
func (d *Dota2) GetChatMemberCount(
	ctx context.Context,
	channelName string,
	channelType protocol.DOTAChatChannelTypeT,
) (*protocol.CMsgDOTAChatGetMemberCountResponse, error) {
	req := &protocol.CMsgDOTAChatGetMemberCount{
		ChannelName: &channelName,
		ChannelType: &channelType,
	}
	resp := &protocol.CMsgDOTAChatGetMemberCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCount),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCountResponse),
		resp,
	)
}

// GetCrawlCavernClaimedRoomCount gets a crawl cavern claimed room count.
// Request ID: k_EMsgClientToGCCavernCrawlGetClaimedRoomCount
// Response ID: k_EMsgClientToGCCavernCrawlGetClaimedRoomCountResponse
// Request type: CMsgClientToGCCavernCrawlGetClaimedRoomCount
// Response type: CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse
func (d *Dota2) GetCrawlCavernClaimedRoomCount(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCount{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlGetClaimedRoomCount),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlGetClaimedRoomCountResponse),
		resp,
	)
}

// GetCurrentPrivateCoachingSession gets a current private coaching session.
// Request ID: k_EMsgClientToGCGetCurrentPrivateCoachingSession
// Response ID: k_EMsgClientToGCGetCurrentPrivateCoachingSessionResponse
// Request type: CMsgClientToGCGetCurrentPrivateCoachingSession
// Response type: CMsgClientToGCGetCurrentPrivateCoachingSessionResponse
func (d *Dota2) GetCurrentPrivateCoachingSession(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetCurrentPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCGetCurrentPrivateCoachingSession{}
	resp := &protocol.CMsgClientToGCGetCurrentPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetCurrentPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetCurrentPrivateCoachingSessionResponse),
		resp,
	)
}

// GetDPCFavorites gets dpc favorites.
// Request ID: k_EMsgClientToGCGetDPCFavorites
// Response ID: k_EMsgClientToGCGetDPCFavoritesResponse
// Request type: CMsgClientToGCGetDPCFavorites
// Response type: CMsgClientToGCGetDPCFavoritesResponse
func (d *Dota2) GetDPCFavorites(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetDPCFavoritesResponse, error) {
	req := &protocol.CMsgClientToGCGetDPCFavorites{}
	resp := &protocol.CMsgClientToGCGetDPCFavoritesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetDPCFavorites),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetDPCFavoritesResponse),
		resp,
	)
}

// GetEventPoints gets event points.
// Request ID: k_EMsgDOTAGetEventPoints
// Response ID: k_EMsgDOTAGetEventPointsResponse
// Request type: CMsgDOTAGetEventPoints
// Response type: CMsgDOTAGetEventPointsResponse
func (d *Dota2) GetEventPoints(
	ctx context.Context,
	eventID uint32,
	accountID uint32,
) (*protocol.CMsgDOTAGetEventPointsResponse, error) {
	req := &protocol.CMsgDOTAGetEventPoints{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAGetEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetEventPointsResponse),
		resp,
	)
}

// GetFavoritePlayers gets favorite players.
// Request ID: k_EMsgClientToGCGetFavoritePlayers
// Response ID: k_EMsgGCToClientGetFavoritePlayersResponse
// Request type: CMsgClientToGCGetFavoritePlayers
// Response type: CMsgGCToClientGetFavoritePlayersResponse
func (d *Dota2) GetFavoritePlayers(
	ctx context.Context,
	paginationKey uint64,
	paginationCount int32,
) (*protocol.CMsgGCToClientGetFavoritePlayersResponse, error) {
	req := &protocol.CMsgClientToGCGetFavoritePlayers{
		PaginationKey:   &paginationKey,
		PaginationCount: &paginationCount,
	}
	resp := &protocol.CMsgGCToClientGetFavoritePlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetFavoritePlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGetFavoritePlayersResponse),
		resp,
	)
}

// GetFilteredPlayers gets filtered players.
// Request ID: k_EMsgClientToGCGetFilteredPlayers
// Response ID: k_EMsgGCToClientGetFilteredPlayersResponse
// Request type: CMsgClientToGCGetFilteredPlayers
// Response type: CMsgGCToClientGetFilteredPlayersResponse
func (d *Dota2) GetFilteredPlayers(
	ctx context.Context,
) (*protocol.CMsgGCToClientGetFilteredPlayersResponse, error) {
	req := &protocol.CMsgClientToGCGetFilteredPlayers{}
	resp := &protocol.CMsgGCToClientGetFilteredPlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetFilteredPlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGetFilteredPlayersResponse),
		resp,
	)
}

// GetGiftPermissions gets gift permissions.
// Request ID: k_EMsgClientToGCGetGiftPermissions
// Response ID: k_EMsgClientToGCGetGiftPermissionsResponse
// Request type: CMsgClientToGCGetGiftPermissions
// Response type: CMsgClientToGCGetGiftPermissionsResponse
func (d *Dota2) GetGiftPermissions(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetGiftPermissionsResponse, error) {
	req := &protocol.CMsgClientToGCGetGiftPermissions{}
	resp := &protocol.CMsgClientToGCGetGiftPermissionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissionsResponse),
		resp,
	)
}

// GetHeroStandings gets hero standings.
// Request ID: k_EMsgGCGetHeroStandings
// Response ID: k_EMsgGCGetHeroStandingsResponse
// Request type: CMsgGCGetHeroStandings
// Response type: CMsgGCGetHeroStandingsResponse
func (d *Dota2) GetHeroStandings(
	ctx context.Context,
) (*protocol.CMsgGCGetHeroStandingsResponse, error) {
	req := &protocol.CMsgGCGetHeroStandings{}
	resp := &protocol.CMsgGCGetHeroStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStandings),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStandingsResponse),
		resp,
	)
}

// GetHeroStatsHistory gets a hero stats history.
// Request ID: k_EMsgGCGetHeroStatsHistory
// Response ID: k_EMsgGCGetHeroStatsHistoryResponse
// Request type: CMsgGCGetHeroStatsHistory
// Response type: CMsgGCGetHeroStatsHistoryResponse
func (d *Dota2) GetHeroStatsHistory(
	ctx context.Context,
	heroID uint32,
) (*protocol.CMsgGCGetHeroStatsHistoryResponse, error) {
	req := &protocol.CMsgGCGetHeroStatsHistory{
		HeroId: &heroID,
	}
	resp := &protocol.CMsgGCGetHeroStatsHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistoryResponse),
		resp,
	)
}

// GetOWMatchDetails gets ow match details.
// Request ID: k_EMsgClientToGCGetOWMatchDetails
// Response ID: k_EMsgClientToGCGetOWMatchDetailsResponse
// Request type: CMsgClientToGCGetOWMatchDetails
// Response type: CMsgClientToGCGetOWMatchDetailsResponse
func (d *Dota2) GetOWMatchDetails(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetOWMatchDetailsResponse, error) {
	req := &protocol.CMsgClientToGCGetOWMatchDetails{}
	resp := &protocol.CMsgClientToGCGetOWMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetOWMatchDetails),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetOWMatchDetailsResponse),
		resp,
	)
}

// GetPeriodicResource gets a periodic resource.
// Request ID: k_EMsgDOTAGetPeriodicResource
// Response ID: k_EMsgDOTAGetPeriodicResourceResponse
// Request type: CMsgDOTAGetPeriodicResource
// Response type: CMsgDOTAGetPeriodicResourceResponse
func (d *Dota2) GetPeriodicResource(
	ctx context.Context,
	accountID uint32,
	periodicResourceID uint32,
	timestamp uint32,
) (*protocol.CMsgDOTAGetPeriodicResourceResponse, error) {
	req := &protocol.CMsgDOTAGetPeriodicResource{
		AccountId:          &accountID,
		PeriodicResourceId: &periodicResourceID,
		Timestamp:          &timestamp,
	}
	resp := &protocol.CMsgDOTAGetPeriodicResourceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResource),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResourceResponse),
		resp,
	)
}

// GetPlayerCardItemInfo gets a player card item info.
// Request ID: k_EMsgGCGetPlayerCardItemInfo
// Response ID: k_EMsgGCGetPlayerCardItemInfoResponse
// Request type: CMsgGCGetPlayerCardItemInfo
// Response type: CMsgGCGetPlayerCardItemInfoResponse
func (d *Dota2) GetPlayerCardItemInfo(
	ctx context.Context,
	accountID uint32,
	playerCardItemIDs []uint64,
	allForEvent uint32,
) (*protocol.CMsgGCGetPlayerCardItemInfoResponse, error) {
	req := &protocol.CMsgGCGetPlayerCardItemInfo{
		AccountId:         &accountID,
		PlayerCardItemIds: playerCardItemIDs,
		AllForEvent:       &allForEvent,
	}
	resp := &protocol.CMsgGCGetPlayerCardItemInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfoResponse),
		resp,
	)
}

// GetPlayerMatchHistory gets a player match history.
// Request ID: k_EMsgDOTAGetPlayerMatchHistory
// Response ID: k_EMsgDOTAGetPlayerMatchHistoryResponse
// Request type: CMsgDOTAGetPlayerMatchHistory
// Response type: CMsgDOTAGetPlayerMatchHistoryResponse
func (d *Dota2) GetPlayerMatchHistory(
	ctx context.Context,
	req *protocol.CMsgDOTAGetPlayerMatchHistory,
) (*protocol.CMsgDOTAGetPlayerMatchHistoryResponse, error) {
	resp := &protocol.CMsgDOTAGetPlayerMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistoryResponse),
		resp,
	)
}

// GetProfileCard gets a profile card.
// Request ID: k_EMsgClientToGCGetProfileCard
// Response ID: k_EMsgClientToGCGetProfileCardResponse
// Request type: CMsgClientToGCGetProfileCard
// Response type: CMsgDOTAProfileCard
func (d *Dota2) GetProfileCard(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgDOTAProfileCard, error) {
	req := &protocol.CMsgClientToGCGetProfileCard{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAProfileCard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse),
		resp,
	)
}

// GetProfileTickets gets profile tickets.
// Request ID: k_EMsgClientToGCGetProfileTickets
// Response ID: k_EMsgClientToGCGetProfileTicketsResponse
// Request type: CMsgClientToGCGetProfileTickets
// Response type: CMsgDOTAProfileTickets
func (d *Dota2) GetProfileTickets(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgDOTAProfileTickets, error) {
	req := &protocol.CMsgClientToGCGetProfileTickets{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAProfileTickets{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileTickets),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse),
		resp,
	)
}

// GetQuestProgress gets quest progress.
// Request ID: k_EMsgClientToGCGetQuestProgress
// Response ID: k_EMsgClientToGCGetQuestProgressResponse
// Request type: CMsgClientToGCGetQuestProgress
// Response type: CMsgClientToGCGetQuestProgressResponse
func (d *Dota2) GetQuestProgress(
	ctx context.Context,
	questIDs []uint32,
) (*protocol.CMsgClientToGCGetQuestProgressResponse, error) {
	req := &protocol.CMsgClientToGCGetQuestProgress{
		QuestIds: questIDs,
	}
	resp := &protocol.CMsgClientToGCGetQuestProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgress),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgressResponse),
		resp,
	)
}

// GetTourneyWeekendPlayerStats gets tourney weekend player stats.
// Request ID: k_EMsgClientToGCWeekendTourneyGetPlayerStats
// Response ID: k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse
// Request type: CMsgDOTAWeekendTourneyPlayerStatsRequest
// Response type: CMsgDOTAWeekendTourneyPlayerStats
func (d *Dota2) GetTourneyWeekendPlayerStats(
	ctx context.Context,
	accountID uint32,
	seasonTrophyID uint32,
) (*protocol.CMsgDOTAWeekendTourneyPlayerStats, error) {
	req := &protocol.CMsgDOTAWeekendTourneyPlayerStatsRequest{
		AccountId:      &accountID,
		SeasonTrophyId: &seasonTrophyID,
	}
	resp := &protocol.CMsgDOTAWeekendTourneyPlayerStats{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse),
		resp,
	)
}

// GetWeekendTourneySchedule gets a weekend tourney schedule.
// Request ID: k_EMsgDOTAGetWeekendTourneySchedule
// Request type: CMsgRequestWeekendTourneySchedule
func (d *Dota2) GetWeekendTourneySchedule() {
	req := &protocol.CMsgRequestWeekendTourneySchedule{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule), req)
}

// GrantDevEventAction grants a dev event action.
// Request ID: k_EMsgDevGrantEventAction
// Response ID: k_EMsgDevGrantEventActionResponse
// Request type: CMsgDevGrantEventAction
// Response type: CMsgDevGrantEventActionResponse
func (d *Dota2) GrantDevEventAction(
	ctx context.Context,
	eventID protocol.EEvent,
	actionID uint32,
	actionScore uint32,
) (*protocol.CMsgDevGrantEventActionResponse, error) {
	req := &protocol.CMsgDevGrantEventAction{
		EventId:     &eventID,
		ActionId:    &actionID,
		ActionScore: &actionScore,
	}
	resp := &protocol.CMsgDevGrantEventActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventActionResponse),
		resp,
	)
}

// GrantDevEventPoints grants dev event points.
// Request ID: k_EMsgDevGrantEventPoints
// Response ID: k_EMsgDevGrantEventPointsResponse
// Request type: CMsgDevGrantEventPoints
// Response type: CMsgDevGrantEventPointsResponse
func (d *Dota2) GrantDevEventPoints(
	ctx context.Context,
	eventID protocol.EEvent,
	eventPoints uint32,
	premiumPoints uint32,
) (*protocol.CMsgDevGrantEventPointsResponse, error) {
	req := &protocol.CMsgDevGrantEventPoints{
		EventId:       &eventID,
		EventPoints:   &eventPoints,
		PremiumPoints: &premiumPoints,
	}
	resp := &protocol.CMsgDevGrantEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventPointsResponse),
		resp,
	)
}

// GrantEventSupportConsumeItem grants a event support consume item.
// Request ID: k_EMsgConsumeEventSupportGrantItem
// Response ID: k_EMsgConsumeEventSupportGrantItemResponse
// Request type: CMsgConsumeEventSupportGrantItem
// Response type: CMsgConsumeEventSupportGrantItemResponse
func (d *Dota2) GrantEventSupportConsumeItem(
	ctx context.Context,
	itemID uint64,
) (*protocol.CMsgConsumeEventSupportGrantItemResponse, error) {
	req := &protocol.CMsgConsumeEventSupportGrantItem{
		ItemId: &itemID,
	}
	resp := &protocol.CMsgConsumeEventSupportGrantItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgConsumeEventSupportGrantItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgConsumeEventSupportGrantItemResponse),
		resp,
	)
}

// InvitePlayerToTeam is undocumented.
// Request ID: k_EMsgGCTeamInvite_InviterToGC
// Response ID: k_EMsgGCTeamInvite_GCImmediateResponseToInviter
// Request type: CMsgDOTATeamInvite_InviterToGC
// Response type: CMsgDOTATeamInvite_GCImmediateResponseToInviter
func (d *Dota2) InvitePlayerToTeam(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTATeamInvite_GCImmediateResponseToInviter, error) {
	req := &protocol.CMsgDOTATeamInvite_InviterToGC{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &protocol.CMsgDOTATeamInvite_GCImmediateResponseToInviter{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter),
		resp,
	)
}

// InvitePrivateChatMember is undocumented.
// Request ID: k_EMsgClientToGCPrivateChatInvite
// Request type: CMsgClientToGCPrivateChatInvite
func (d *Dota2) InvitePrivateChatMember(
	privateChatChannelName string,
	invitedAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatInvite{
		PrivateChatChannelName: &privateChatChannelName,
		InvitedAccountId:       &invitedAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite), req)
}

// JoinChatChannel joins a chat channel.
// Request ID: k_EMsgGCJoinChatChannel
// Response ID: k_EMsgGCJoinChatChannelResponse
// Request type: CMsgDOTAJoinChatChannel
// Response type: CMsgDOTAJoinChatChannelResponse
func (d *Dota2) JoinChatChannel(
	ctx context.Context,
	channelName string,
	channelType protocol.DOTAChatChannelTypeT,
	silentRejection bool,
) (*protocol.CMsgDOTAJoinChatChannelResponse, error) {
	req := &protocol.CMsgDOTAJoinChatChannel{
		ChannelName:     &channelName,
		ChannelType:     &channelType,
		SilentRejection: &silentRejection,
	}
	resp := &protocol.CMsgDOTAJoinChatChannelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinChatChannel),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse),
		resp,
	)
}

// JoinGuild joins a guild.
// Request ID: k_EMsgClientToGCJoinGuild
// Response ID: k_EMsgClientToGCJoinGuildResponse
// Request type: CMsgClientToGCJoinGuild
// Response type: CMsgClientToGCJoinGuildResponse
func (d *Dota2) JoinGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCJoinGuildResponse, error) {
	req := &protocol.CMsgClientToGCJoinGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCJoinGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinGuildResponse),
		resp,
	)
}

// JoinLobby joins a lobby.
// Request ID: k_EMsgGCPracticeLobbyJoin
// Response ID: k_EMsgGCPracticeLobbyJoinResponse
// Request type: CMsgPracticeLobbyJoin
// Response type: CMsgPracticeLobbyJoinResponse
func (d *Dota2) JoinLobby(
	ctx context.Context,
	lobbyID uint64,
	passKey string,
	customGameCrc uint64,
	customGameTimestamp uint32,
) (*protocol.CMsgPracticeLobbyJoinResponse, error) {
	req := &protocol.CMsgPracticeLobbyJoin{
		LobbyId:             &lobbyID,
		PassKey:             &passKey,
		CustomGameCrc:       &customGameCrc,
		CustomGameTimestamp: &customGameTimestamp,
	}
	resp := &protocol.CMsgPracticeLobbyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoin),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinResponse),
		resp,
	)
}

// JoinLobbyBroadcastChannel joins a lobby broadcast channel.
// Request ID: k_EMsgGCPracticeLobbyJoinBroadcastChannel
// Request type: CMsgPracticeLobbyJoinBroadcastChannel
func (d *Dota2) JoinLobbyBroadcastChannel(
	channel uint32,
	preferredDescription string,
	preferredCountryCode string,
	preferredLanguageCode string,
) {
	req := &protocol.CMsgPracticeLobbyJoinBroadcastChannel{
		Channel:               &channel,
		PreferredDescription:  &preferredDescription,
		PreferredCountryCode:  &preferredCountryCode,
		PreferredLanguageCode: &preferredLanguageCode,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinBroadcastChannel), req)
}

// JoinPartyFromBeacon joins a party from beacon.
// Request ID: k_EMsgClientToGCJoinPartyFromBeacon
// Response ID: k_EMsgGCToClientJoinPartyFromBeaconResponse
// Request type: CMsgClientToGCJoinPartyFromBeacon
// Response type: CMsgGCToClientJoinPartyFromBeaconResponse
func (d *Dota2) JoinPartyFromBeacon(
	ctx context.Context,
	partyID uint64,
	accountID uint32,
	beaconType int32,
) (*protocol.CMsgGCToClientJoinPartyFromBeaconResponse, error) {
	req := &protocol.CMsgClientToGCJoinPartyFromBeacon{
		PartyId:    &partyID,
		AccountId:  &accountID,
		BeaconType: &beaconType,
	}
	resp := &protocol.CMsgGCToClientJoinPartyFromBeaconResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPartyFromBeacon),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientJoinPartyFromBeaconResponse),
		resp,
	)
}

// JoinPlaytest joins a playtest.
// Request ID: k_EMsgClientToGCJoinPlaytest
// Response ID: k_EMsgClientToGCJoinPlaytestResponse
// Request type: CMsgClientToGCJoinPlaytest
// Response type: CMsgClientToGCJoinPlaytestResponse
func (d *Dota2) JoinPlaytest(
	ctx context.Context,
) (*protocol.CMsgClientToGCJoinPlaytestResponse, error) {
	req := &protocol.CMsgClientToGCJoinPlaytest{}
	resp := &protocol.CMsgClientToGCJoinPlaytestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytestResponse),
		resp,
	)
}

// JoinPrivateCoachingSessionLobby joins a private coaching session lobby.
// Request ID: k_EMsgClientToGCJoinPrivateCoachingSessionLobby
// Response ID: k_EMsgClientToGCJoinPrivateCoachingSessionLobbyResponse
// Request type: CMsgClientToGCJoinPrivateCoachingSessionLobby
// Response type: CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse
func (d *Dota2) JoinPrivateCoachingSessionLobby(
	ctx context.Context,
) (*protocol.CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse, error) {
	req := &protocol.CMsgClientToGCJoinPrivateCoachingSessionLobby{}
	resp := &protocol.CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPrivateCoachingSessionLobby),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPrivateCoachingSessionLobbyResponse),
		resp,
	)
}

// JoinQuickCustomLobby joins a quick custom lobby.
// Request ID: k_EMsgGCQuickJoinCustomLobby
// Response ID: k_EMsgGCQuickJoinCustomLobbyResponse
// Request type: CMsgQuickJoinCustomLobby
// Response type: CMsgQuickJoinCustomLobbyResponse
func (d *Dota2) JoinQuickCustomLobby(
	ctx context.Context,
	req *protocol.CMsgQuickJoinCustomLobby,
) (*protocol.CMsgQuickJoinCustomLobbyResponse, error) {
	resp := &protocol.CMsgQuickJoinCustomLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobby),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobbyResponse),
		resp,
	)
}

// KickGuildMember kicks a guild member.
// Request ID: k_EMsgClientToGCKickGuildMember
// Response ID: k_EMsgClientToGCKickGuildMemberResponse
// Request type: CMsgClientToGCKickGuildMember
// Response type: CMsgClientToGCKickGuildMemberResponse
func (d *Dota2) KickGuildMember(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCKickGuildMemberResponse, error) {
	req := &protocol.CMsgClientToGCKickGuildMember{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCKickGuildMemberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCKickGuildMember),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCKickGuildMemberResponse),
		resp,
	)
}

// KickLobbyMember kicks a lobby member.
// Request ID: k_EMsgGCPracticeLobbyKick
// Request type: CMsgPracticeLobbyKick
func (d *Dota2) KickLobbyMember(
	accountID uint32,
) {
	req := &protocol.CMsgPracticeLobbyKick{
		AccountId: &accountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick), req)
}

// KickLobbyMemberFromTeam kicks a lobby member from team.
// Request ID: k_EMsgGCPracticeLobbyKickFromTeam
// Request type: CMsgPracticeLobbyKickFromTeam
func (d *Dota2) KickLobbyMemberFromTeam(
	accountID uint32,
) {
	req := &protocol.CMsgPracticeLobbyKickFromTeam{
		AccountId: &accountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam), req)
}

// KickPrivateChatMember kicks a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatKick
// Request type: CMsgClientToGCPrivateChatKick
func (d *Dota2) KickPrivateChatMember(
	privateChatChannelName string,
	kickAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatKick{
		PrivateChatChannelName: &privateChatChannelName,
		KickAccountId:          &kickAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick), req)
}

// KickTeamMember kicks a team member.
// Request ID: k_EMsgGCKickTeamMember
// Response ID: k_EMsgGCKickTeamMemberResponse
// Request type: CMsgDOTAKickTeamMember
// Response type: CMsgDOTAKickTeamMemberResponse
func (d *Dota2) KickTeamMember(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTAKickTeamMemberResponse, error) {
	req := &protocol.CMsgDOTAKickTeamMember{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &protocol.CMsgDOTAKickTeamMemberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCKickTeamMember),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCKickTeamMemberResponse),
		resp,
	)
}

// LaunchLobby launchs a lobby.
// Request ID: k_EMsgGCPracticeLobbyLaunch
// Request type: CMsgPracticeLobbyLaunch
func (d *Dota2) LaunchLobby() {
	req := &protocol.CMsgPracticeLobbyLaunch{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyLaunch), req)
}

// LeaveChatChannel leaves a chat channel.
// Request ID: k_EMsgGCLeaveChatChannel
// Request type: CMsgDOTALeaveChatChannel
func (d *Dota2) LeaveChatChannel(
	channelID uint64,
) {
	req := &protocol.CMsgDOTALeaveChatChannel{
		ChannelId: &channelID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveChatChannel), req)
}

// LeaveGuild leaves a guild.
// Request ID: k_EMsgClientToGCLeaveGuild
// Response ID: k_EMsgClientToGCLeaveGuildResponse
// Request type: CMsgClientToGCLeaveGuild
// Response type: CMsgClientToGCLeaveGuildResponse
func (d *Dota2) LeaveGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCLeaveGuildResponse, error) {
	req := &protocol.CMsgClientToGCLeaveGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCLeaveGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeaveGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeaveGuildResponse),
		resp,
	)
}

// LeaveLobby leaves a lobby.
// Request ID: k_EMsgGCPracticeLobbyLeave
// Request type: CMsgPracticeLobbyLeave
func (d *Dota2) LeaveLobby() {
	req := &protocol.CMsgPracticeLobbyLeave{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyLeave), req)
}

// LeavePrivateCoachingSession leaves a private coaching session.
// Request ID: k_EMsgClientToGCLeavePrivateCoachingSession
// Response ID: k_EMsgClientToGCLeavePrivateCoachingSessionResponse
// Request type: CMsgClientToGCLeavePrivateCoachingSession
// Response type: CMsgClientToGCLeavePrivateCoachingSessionResponse
func (d *Dota2) LeavePrivateCoachingSession(
	ctx context.Context,
) (*protocol.CMsgClientToGCLeavePrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCLeavePrivateCoachingSession{}
	resp := &protocol.CMsgClientToGCLeavePrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeavePrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeavePrivateCoachingSessionResponse),
		resp,
	)
}

// LeaveTeam leaves a team.
// Request ID: k_EMsgGCLeaveTeam
// Response ID: k_EMsgGCLeaveTeamResponse
// Request type: CMsgDOTALeaveTeam
// Response type: CMsgDOTALeaveTeamResponse
func (d *Dota2) LeaveTeam(
	ctx context.Context,
	teamID uint32,
) (*protocol.CMsgDOTALeaveTeamResponse, error) {
	req := &protocol.CMsgDOTALeaveTeam{
		TeamId: &teamID,
	}
	resp := &protocol.CMsgDOTALeaveTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveTeam),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveTeamResponse),
		resp,
	)
}

// LeaveTourneyWeekend leaves a tourney weekend.
// Request ID: k_EMsgClientToGCWeekendTourneyLeave
// Request type: CMsgWeekendTourneyLeave
func (d *Dota2) LeaveTourneyWeekend() {
	req := &protocol.CMsgWeekendTourneyLeave{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave), req)
}

// ListChatChannel lists a chat channel.
// Request ID: k_EMsgGCRequestChatChannelList
// Response ID: k_EMsgGCRequestChatChannelListResponse
// Request type: CMsgDOTARequestChatChannelList
// Response type: CMsgDOTARequestChatChannelListResponse
func (d *Dota2) ListChatChannel(
	ctx context.Context,
) (*protocol.CMsgDOTARequestChatChannelListResponse, error) {
	req := &protocol.CMsgDOTARequestChatChannelList{}
	resp := &protocol.CMsgDOTARequestChatChannelListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRequestChatChannelList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRequestChatChannelListResponse),
		resp,
	)
}

// ListCustomGamesTop lists a custom games top.
// Request ID: k_EMsgGCTopCustomGamesList
// Request type: CMsgGCTopCustomGamesList
func (d *Dota2) ListCustomGamesTop(
	topCustomGames []uint64,
	gameOfTheDay uint64,
) {
	req := &protocol.CMsgGCTopCustomGamesList{
		TopCustomGames: topCustomGames,
		GameOfTheDay:   &gameOfTheDay,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTopCustomGamesList), req)
}

// ListFriendLobby lists a friend lobby.
// Request ID: k_EMsgGCFriendPracticeLobbyListRequest
// Response ID: k_EMsgGCFriendPracticeLobbyListResponse
// Request type: CMsgFriendPracticeLobbyListRequest
// Response type: CMsgFriendPracticeLobbyListResponse
func (d *Dota2) ListFriendLobby(
	ctx context.Context,
	friends []uint32,
) (*protocol.CMsgFriendPracticeLobbyListResponse, error) {
	req := &protocol.CMsgFriendPracticeLobbyListRequest{
		Friends: friends,
	}
	resp := &protocol.CMsgFriendPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListResponse),
		resp,
	)
}

// ListLobbies lists lobbies.
// Request ID: k_EMsgGCPracticeLobbyList
// Response ID: k_EMsgGCPracticeLobbyListResponse
// Request type: CMsgPracticeLobbyList
// Response type: CMsgPracticeLobbyListResponse
func (d *Dota2) ListLobbies(
	ctx context.Context,
	passKey string,
	region uint32,
	gameMode protocol.DOTA_GameMode,
) (*protocol.CMsgPracticeLobbyListResponse, error) {
	req := &protocol.CMsgPracticeLobbyList{
		PassKey:  &passKey,
		Region:   &region,
		GameMode: &gameMode,
	}
	resp := &protocol.CMsgPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyListResponse),
		resp,
	)
}

// ListLobbySpectator lists a lobby spectator.
// Request ID: k_EMsgClientToGCSpectatorLobbyList
// Response ID: k_EMsgClientToGCSpectatorLobbyListResponse
// Request type: CMsgSpectatorLobbyList
// Response type: CMsgSpectatorLobbyListResponse
func (d *Dota2) ListLobbySpectator(
	ctx context.Context,
) (*protocol.CMsgSpectatorLobbyListResponse, error) {
	req := &protocol.CMsgSpectatorLobbyList{}
	resp := &protocol.CMsgSpectatorLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyListResponse),
		resp,
	)
}

// ListTrophies lists trophies.
// Request ID: k_EMsgClientToGCGetTrophyList
// Response ID: k_EMsgClientToGCGetTrophyListResponse
// Request type: CMsgClientToGCGetTrophyList
// Response type: CMsgClientToGCGetTrophyListResponse
func (d *Dota2) ListTrophies(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetTrophyListResponse, error) {
	req := &protocol.CMsgClientToGCGetTrophyList{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetTrophyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetTrophyListResponse),
		resp,
	)
}

// OpenPlayerCardPack opens a player card pack.
// Request ID: k_EMsgClientToGCOpenPlayerCardPack
// Response ID: k_EMsgClientToGCOpenPlayerCardPackResponse
// Request type: CMsgClientToGCOpenPlayerCardPack
// Response type: CMsgClientToGCOpenPlayerCardPackResponse
func (d *Dota2) OpenPlayerCardPack(
	ctx context.Context,
	playerCardPackItemID uint64,
	teamID uint32,
	leagueID uint32,
) (*protocol.CMsgClientToGCOpenPlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCOpenPlayerCardPack{
		PlayerCardPackItemId: &playerCardPackItemID,
		TeamId:               &teamID,
		LeagueId:             &leagueID,
	}
	resp := &protocol.CMsgClientToGCOpenPlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPackResponse),
		resp,
	)
}

// PromotePrivateChatMember promotes a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatPromote
// Request type: CMsgClientToGCPrivateChatPromote
func (d *Dota2) PromotePrivateChatMember(
	privateChatChannelName string,
	promoteAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatPromote{
		PrivateChatChannelName: &privateChatChannelName,
		PromoteAccountId:       &promoteAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote), req)
}

// PublishUserStat publishs a user stat.
// Request ID: k_EMsgClientToGCPublishUserStat
// Request type: CMsgClientToGCPublishUserStat
func (d *Dota2) PublishUserStat(
	userStatsEvent uint32,
	referenceData uint64,
) {
	req := &protocol.CMsgClientToGCPublishUserStat{
		UserStatsEvent: &userStatsEvent,
		ReferenceData:  &referenceData,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPublishUserStat), req)
}

// PurchaseHeroRandomRelic purchases a hero random relic.
// Request ID: k_EMsgPurchaseHeroRandomRelic
// Response ID: k_EMsgPurchaseHeroRandomRelicResponse
// Request type: CMsgPurchaseHeroRandomRelic
// Response type: CMsgPurchaseHeroRandomRelicResponse
func (d *Dota2) PurchaseHeroRandomRelic(
	ctx context.Context,
	heroID uint32,
	relicRarity protocol.EHeroRelicRarity,
) (*protocol.CMsgPurchaseHeroRandomRelicResponse, error) {
	req := &protocol.CMsgPurchaseHeroRandomRelic{
		HeroId:      &heroID,
		RelicRarity: &relicRarity,
	}
	resp := &protocol.CMsgPurchaseHeroRandomRelicResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseHeroRandomRelic),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseHeroRandomRelicResponse),
		resp,
	)
}

// PurchaseItemWithEventPoints purchases item with event points.
// Request ID: k_EMsgPurchaseItemWithEventPoints
// Response ID: k_EMsgPurchaseItemWithEventPointsResponse
// Request type: CMsgPurchaseItemWithEventPoints
// Response type: CMsgPurchaseItemWithEventPointsResponse
func (d *Dota2) PurchaseItemWithEventPoints(
	ctx context.Context,
	itemDef uint32,
	quantity uint32,
	eventID protocol.EEvent,
	usePremiumPoints bool,
) (*protocol.CMsgPurchaseItemWithEventPointsResponse, error) {
	req := &protocol.CMsgPurchaseItemWithEventPoints{
		ItemDef:          &itemDef,
		Quantity:         &quantity,
		EventId:          &eventID,
		UsePremiumPoints: &usePremiumPoints,
	}
	resp := &protocol.CMsgPurchaseItemWithEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPointsResponse),
		resp,
	)
}

// PurchaseLabyrinthBlessings purchases labyrinth blessings.
// Request ID: k_EMsgClientToGCPurchaseLabyrinthBlessings
// Response ID: k_EMsgClientToGCPurchaseLabyrinthBlessingsResponse
// Request type: CMsgClientToGCPurchaseLabyrinthBlessings
// Response type: CMsgClientToGCPurchaseLabyrinthBlessingsResponse
func (d *Dota2) PurchaseLabyrinthBlessings(
	ctx context.Context,
	eventID protocol.EEvent,
	blessingIDs []int32,
	debug bool,
	debugRemove bool,
) (*protocol.CMsgClientToGCPurchaseLabyrinthBlessingsResponse, error) {
	req := &protocol.CMsgClientToGCPurchaseLabyrinthBlessings{
		EventId:     &eventID,
		BlessingIds: blessingIDs,
		Debug:       &debug,
		DebugRemove: &debugRemove,
	}
	resp := &protocol.CMsgClientToGCPurchaseLabyrinthBlessingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPurchaseLabyrinthBlessings),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPurchaseLabyrinthBlessingsResponse),
		resp,
	)
}

// PurchasePlayerCardSpecific purchases a player card specific.
// Request ID: k_EMsgClientToGCPlayerCardSpecificPurchaseRequest
// Response ID: k_EMsgClientToGCPlayerCardSpecificPurchaseResponse
// Request type: CMsgClientToGCPlayerCardSpecificPurchaseRequest
// Response type: CMsgClientToGCPlayerCardSpecificPurchaseResponse
func (d *Dota2) PurchasePlayerCardSpecific(
	ctx context.Context,
	playerAccountID uint32,
	eventID uint32,
	cardDustItemID uint64,
) (*protocol.CMsgClientToGCPlayerCardSpecificPurchaseResponse, error) {
	req := &protocol.CMsgClientToGCPlayerCardSpecificPurchaseRequest{
		PlayerAccountId: &playerAccountID,
		EventId:         &eventID,
		CardDustItemId:  &cardDustItemID,
	}
	resp := &protocol.CMsgClientToGCPlayerCardSpecificPurchaseResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerCardSpecificPurchaseRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerCardSpecificPurchaseResponse),
		resp,
	)
}

// QueryHasItem queries to check if the target has item.
// Request ID: k_EMsgGCHasItemQuery
// Request type: CMsgDOTAHasItemQuery
func (d *Dota2) QueryHasItem(
	accountID uint32,
	itemID uint64,
) {
	req := &protocol.CMsgDOTAHasItemQuery{
		AccountId: &accountID,
		ItemId:    &itemID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCHasItemQuery), req)
}

// RecordContestVote records a contest vote.
// Request ID: k_EMsgClientToGCRecordContestVote
// Response ID: k_EMsgGCToClientRecordContestVoteResponse
// Request type: CMsgClientToGCRecordContestVote
// Response type: CMsgGCToClientRecordContestVoteResponse
func (d *Dota2) RecordContestVote(
	ctx context.Context,
	contestID uint32,
	contestItemID uint64,
	vote int32,
) (*protocol.CMsgGCToClientRecordContestVoteResponse, error) {
	req := &protocol.CMsgClientToGCRecordContestVote{
		ContestId:     &contestID,
		ContestItemId: &contestItemID,
		Vote:          &vote,
	}
	resp := &protocol.CMsgGCToClientRecordContestVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecordContestVote),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRecordContestVoteResponse),
		resp,
	)
}

// RecyclePlayerCard recycles a player card.
// Request ID: k_EMsgClientToGCRecyclePlayerCard
// Response ID: k_EMsgClientToGCRecyclePlayerCardResponse
// Request type: CMsgClientToGCRecyclePlayerCard
// Response type: CMsgClientToGCRecyclePlayerCardResponse
func (d *Dota2) RecyclePlayerCard(
	ctx context.Context,
	playerCardItemIDs []uint64,
	eventID uint32,
) (*protocol.CMsgClientToGCRecyclePlayerCardResponse, error) {
	req := &protocol.CMsgClientToGCRecyclePlayerCard{
		PlayerCardItemIds: playerCardItemIDs,
		EventId:           &eventID,
	}
	resp := &protocol.CMsgClientToGCRecyclePlayerCardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCardResponse),
		resp,
	)
}

// RedeemDraftUnderReward redeems a draft under reward.
// Request ID: k_EMsgClientToGCUnderDraftRedeemReward
// Response ID: k_EMsgClientToGCUnderDraftRedeemRewardResponse
// Request type: CMsgClientToGCUnderDraftRedeemReward
// Response type: CMsgClientToGCUnderDraftRedeemRewardResponse
func (d *Dota2) RedeemDraftUnderReward(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
) (*protocol.CMsgClientToGCUnderDraftRedeemRewardResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRedeemReward{
		EventId:  &eventID,
		ActionId: &actionID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRedeemRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRedeemReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRedeemRewardResponse),
		resp,
	)
}

// RedeemItem redeems a item.
// Request ID: k_EMsgDOTARedeemItem
// Response ID: k_EMsgDOTARedeemItemResponse
// Request type: CMsgDOTARedeemItem
// Response type: CMsgDOTARedeemItemResponse
func (d *Dota2) RedeemItem(
	ctx context.Context,
	currencyID uint64,
	purchaseDef uint32,
) (*protocol.CMsgDOTARedeemItemResponse, error) {
	req := &protocol.CMsgDOTARedeemItem{
		CurrencyId:  &currencyID,
		PurchaseDef: &purchaseDef,
	}
	resp := &protocol.CMsgDOTARedeemItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTARedeemItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTARedeemItemResponse),
		resp,
	)
}

// RejoinAllChatChannels is undocumented.
// Request ID: k_EMsgClientsRejoinChatChannels
// Request type: CMsgClientsRejoinChatChannels
func (d *Dota2) RejoinAllChatChannels() {
	req := &protocol.CMsgClientsRejoinChatChannels{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels), req)
}

// ReleaseEditorItemReservation releases a editor item reservation.
// Request ID: k_EMsgGCItemEditorReleaseReservation
// Response ID: k_EMsgGCItemEditorReleaseReservationResponse
// Request type: CMsgGCItemEditorReleaseReservation
// Response type: CMsgGCItemEditorReleaseReservationResponse
func (d *Dota2) ReleaseEditorItemReservation(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*protocol.CMsgGCItemEditorReleaseReservationResponse, error) {
	req := &protocol.CMsgGCItemEditorReleaseReservation{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &protocol.CMsgGCItemEditorReleaseReservationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservation),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservationResponse),
		resp,
	)
}

// ReportChatPublicSpam reports a chat public spam.
// Request ID: k_EMsgGCChatReportPublicSpam
// Request type: CMsgGCChatReportPublicSpam
func (d *Dota2) ReportChatPublicSpam(
	channelID uint64,
	channelUserID uint32,
) {
	req := &protocol.CMsgGCChatReportPublicSpam{
		ChannelId:     &channelID,
		ChannelUserId: &channelUserID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCChatReportPublicSpam), req)
}

// ReportGuildContent reports a guild content.
// Request ID: k_EMsgClientToGCReportGuildContent
// Response ID: k_EMsgClientToGCReportGuildContentResponse
// Request type: CMsgClientToGCReportGuildContent
// Response type: CMsgClientToGCReportGuildContentResponse
func (d *Dota2) ReportGuildContent(
	ctx context.Context,
	guildID uint32,
	guildContentFlags uint32,
) (*protocol.CMsgClientToGCReportGuildContentResponse, error) {
	req := &protocol.CMsgClientToGCReportGuildContent{
		GuildId:           &guildID,
		GuildContentFlags: &guildContentFlags,
	}
	resp := &protocol.CMsgClientToGCReportGuildContentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCReportGuildContent),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCReportGuildContentResponse),
		resp,
	)
}

// RequestAccountGuildEventData requests a account guild event data.
// Request ID: k_EMsgClientToGCRequestAccountGuildEventData
// Response ID: k_EMsgClientToGCRequestAccountGuildEventDataResponse
// Request type: CMsgClientToGCRequestAccountGuildEventData
// Response type: CMsgClientToGCRequestAccountGuildEventDataResponse
func (d *Dota2) RequestAccountGuildEventData(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestAccountGuildEventDataResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildEventData{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildEventDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildEventData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildEventDataResponse),
		resp,
	)
}

// RequestAccountGuildPersonaInfo requests a account guild persona info.
// Request ID: k_EMsgClientToGCRequestAccountGuildPersonaInfo
// Response ID: k_EMsgClientToGCRequestAccountGuildPersonaInfoResponse
// Request type: CMsgClientToGCRequestAccountGuildPersonaInfo
// Response type: CMsgClientToGCRequestAccountGuildPersonaInfoResponse
func (d *Dota2) RequestAccountGuildPersonaInfo(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCRequestAccountGuildPersonaInfoResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfo{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoResponse),
		resp,
	)
}

// RequestAccountGuildPersonaInfoBatch requests a account guild persona info batch.
// Request ID: k_EMsgClientToGCRequestAccountGuildPersonaInfoBatch
// Response ID: k_EMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse
// Request type: CMsgClientToGCRequestAccountGuildPersonaInfoBatch
// Response type: CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse
func (d *Dota2) RequestAccountGuildPersonaInfoBatch(
	ctx context.Context,
	accountIDs []uint32,
) (*protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatch{
		AccountIds: accountIDs,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoBatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse),
		resp,
	)
}

// RequestActiveBeaconParties requests active beacon parties.
// Request ID: k_EMsgClientToGCRequestActiveBeaconParties
// Response ID: k_EMsgGCToClientRequestActiveBeaconPartiesResponse
// Request type: CMsgClientToGCRequestActiveBeaconParties
// Response type: CMsgGCToClientRequestActiveBeaconPartiesResponse
func (d *Dota2) RequestActiveBeaconParties(
	ctx context.Context,
) (*protocol.CMsgGCToClientRequestActiveBeaconPartiesResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveBeaconParties{}
	resp := &protocol.CMsgGCToClientRequestActiveBeaconPartiesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveBeaconParties),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestActiveBeaconPartiesResponse),
		resp,
	)
}

// RequestActiveGuildChallenge requests a active guild challenge.
// Request ID: k_EMsgClientToGCRequestActiveGuildChallenge
// Response ID: k_EMsgClientToGCRequestActiveGuildChallengeResponse
// Request type: CMsgClientToGCRequestActiveGuildChallenge
// Response type: CMsgClientToGCRequestActiveGuildChallengeResponse
func (d *Dota2) RequestActiveGuildChallenge(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestActiveGuildChallengeResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveGuildChallenge{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestActiveGuildChallengeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildChallenge),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildChallengeResponse),
		resp,
	)
}

// RequestActiveGuildContracts requests active guild contracts.
// Request ID: k_EMsgClientToGCRequestActiveGuildContracts
// Response ID: k_EMsgClientToGCRequestActiveGuildContractsResponse
// Request type: CMsgClientToGCRequestActiveGuildContracts
// Response type: CMsgClientToGCRequestActiveGuildContractsResponse
func (d *Dota2) RequestActiveGuildContracts(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestActiveGuildContractsResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveGuildContracts{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestActiveGuildContractsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildContracts),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildContractsResponse),
		resp,
	)
}

// RequestAnchorPhoneNumber requests to check if the target anchor phone number.
// Request ID: k_EMsgAnchorPhoneNumberRequest
// Response ID: k_EMsgAnchorPhoneNumberResponse
// Request type: CMsgDOTAAnchorPhoneNumberRequest
// Response type: CMsgDOTAAnchorPhoneNumberResponse
func (d *Dota2) RequestAnchorPhoneNumber(
	ctx context.Context,
) (*protocol.CMsgDOTAAnchorPhoneNumberResponse, error) {
	req := &protocol.CMsgDOTAAnchorPhoneNumberRequest{}
	resp := &protocol.CMsgDOTAAnchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgAnchorPhoneNumberRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgAnchorPhoneNumberResponse),
		resp,
	)
}

// RequestArcanaVotesRemaining requests a arcana votes remaining.
// Request ID: k_EMsgClientToGCRequestArcanaVotesRemaining
// Response ID: k_EMsgClientToGCRequestArcanaVotesRemainingResponse
// Request type: CMsgClientToGCRequestArcanaVotesRemaining
// Response type: CMsgClientToGCRequestArcanaVotesRemainingResponse
func (d *Dota2) RequestArcanaVotesRemaining(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestArcanaVotesRemainingResponse, error) {
	req := &protocol.CMsgClientToGCRequestArcanaVotesRemaining{}
	resp := &protocol.CMsgClientToGCRequestArcanaVotesRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemaining),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemainingResponse),
		resp,
	)
}

// RequestChinaSSAAccepted requests a china ssa accepted.
// Request ID: k_EMsgClientToGCChinaSSAAcceptedRequest
// Response ID: k_EMsgClientToGCChinaSSAAcceptedResponse
// Request type: CMsgClientToGCChinaSSAAcceptedRequest
// Response type: CMsgClientToGCChinaSSAAcceptedResponse
func (d *Dota2) RequestChinaSSAAccepted(
	ctx context.Context,
) (*protocol.CMsgClientToGCChinaSSAAcceptedResponse, error) {
	req := &protocol.CMsgClientToGCChinaSSAAcceptedRequest{}
	resp := &protocol.CMsgClientToGCChinaSSAAcceptedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAAcceptedRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAAcceptedResponse),
		resp,
	)
}

// RequestChinaSSAURL requests a china ssaurl.
// Request ID: k_EMsgClientToGCChinaSSAURLRequest
// Response ID: k_EMsgClientToGCChinaSSAURLResponse
// Request type: CMsgClientToGCChinaSSAURLRequest
// Response type: CMsgClientToGCChinaSSAURLResponse
func (d *Dota2) RequestChinaSSAURL(
	ctx context.Context,
) (*protocol.CMsgClientToGCChinaSSAURLResponse, error) {
	req := &protocol.CMsgClientToGCChinaSSAURLRequest{}
	resp := &protocol.CMsgClientToGCChinaSSAURLResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAURLRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAURLResponse),
		resp,
	)
}

// RequestCompendiumData requests a compendium data.
// Request ID: k_EMsgGCCompendiumDataRequest
// Response ID: k_EMsgGCCompendiumDataResponse
// Request type: CMsgDOTACompendiumDataRequest
// Response type: CMsgDOTACompendiumDataResponse
func (d *Dota2) RequestCompendiumData(
	ctx context.Context,
	accountID uint32,
	leagueid uint32,
) (*protocol.CMsgDOTACompendiumDataResponse, error) {
	req := &protocol.CMsgDOTACompendiumDataRequest{
		AccountId: &accountID,
		Leagueid:  &leagueid,
	}
	resp := &protocol.CMsgDOTACompendiumDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumDataResponse),
		resp,
	)
}

// RequestContestVotes requests contest votes.
// Request ID: k_EMsgClientToGCRequestContestVotes
// Response ID: k_EMsgClientToGCRequestContestVotesResponse
// Request type: CMsgClientToGCRequestContestVotes
// Response type: CMsgClientToGCRequestContestVotesResponse
func (d *Dota2) RequestContestVotes(
	ctx context.Context,
	contestID uint32,
) (*protocol.CMsgClientToGCRequestContestVotesResponse, error) {
	req := &protocol.CMsgClientToGCRequestContestVotes{
		ContestId: &contestID,
	}
	resp := &protocol.CMsgClientToGCRequestContestVotesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestContestVotes),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestContestVotesResponse),
		resp,
	)
}

// RequestCrawlCavernMapState requests a crawl cavern map state.
// Request ID: k_EMsgClientToGCCavernCrawlRequestMapState
// Response ID: k_EMsgClientToGCCavernCrawlRequestMapStateResponse
// Request type: CMsgClientToGCCavernCrawlRequestMapState
// Response type: CMsgClientToGCCavernCrawlRequestMapStateResponse
func (d *Dota2) RequestCrawlCavernMapState(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCCavernCrawlRequestMapStateResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlRequestMapState{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlRequestMapStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlRequestMapState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlRequestMapStateResponse),
		resp,
	)
}

// RequestCustomGamesFriendsPlayed requests a custom games friends played.
// Request ID: k_EMsgClientToGCCustomGamesFriendsPlayedRequest
// Response ID: k_EMsgGCToClientCustomGamesFriendsPlayedResponse
// Request type: CMsgClientToGCCustomGamesFriendsPlayedRequest
// Response type: CMsgGCToClientCustomGamesFriendsPlayedResponse
func (d *Dota2) RequestCustomGamesFriendsPlayed(
	ctx context.Context,
) (*protocol.CMsgGCToClientCustomGamesFriendsPlayedResponse, error) {
	req := &protocol.CMsgClientToGCCustomGamesFriendsPlayedRequest{}
	resp := &protocol.CMsgGCToClientCustomGamesFriendsPlayedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCustomGamesFriendsPlayedRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCustomGamesFriendsPlayedResponse),
		resp,
	)
}

// RequestEmoticonData requests a emoticon data.
// Request ID: k_EMsgClientToGCEmoticonDataRequest
// Response ID: k_EMsgGCToClientEmoticonData
// Request type: CMsgClientToGCEmoticonDataRequest
// Response type: CMsgGCToClientEmoticonData
func (d *Dota2) RequestEmoticonData(
	ctx context.Context,
) (*protocol.CMsgGCToClientEmoticonData, error) {
	req := &protocol.CMsgClientToGCEmoticonDataRequest{}
	resp := &protocol.CMsgGCToClientEmoticonData{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEmoticonDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientEmoticonData),
		resp,
	)
}

// RequestEventGoals requests event goals.
// Request ID: k_EMsgClientToGCEventGoalsRequest
// Response ID: k_EMsgClientToGCEventGoalsResponse
// Request type: CMsgClientToGCGetEventGoals
// Response type: CMsgEventGoals
func (d *Dota2) RequestEventGoals(
	ctx context.Context,
	eventIDs []protocol.EEvent,
) (*protocol.CMsgEventGoals, error) {
	req := &protocol.CMsgClientToGCGetEventGoals{
		EventIds: eventIDs,
	}
	resp := &protocol.CMsgEventGoals{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse),
		resp,
	)
}

// RequestEventPointLogResponseV2 requests a event point log response v 2.
// Request ID: k_EMsgClientToGCRequestEventPointLogResponseV2
// Request type: CMsgClientToGCRequestEventPointLogResponseV2
func (d *Dota2) RequestEventPointLogResponseV2(
	result bool,
	eventID protocol.EEvent,
	logEntries []*protocol.CMsgClientToGCRequestEventPointLogResponseV2_LogEntry,
) {
	req := &protocol.CMsgClientToGCRequestEventPointLogResponseV2{
		Result:     &result,
		EventId:    &eventID,
		LogEntries: logEntries,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLogResponseV2), req)
}

// RequestEventPointLogV2 requests a event point log v 2.
// Request ID: k_EMsgClientToGCRequestEventPointLogV2
// Request type: CMsgClientToGCRequestEventPointLogV2
func (d *Dota2) RequestEventPointLogV2(
	eventID uint32,
) {
	req := &protocol.CMsgClientToGCRequestEventPointLogV2{
		EventId: &eventID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLogV2), req)
}

// RequestEventTipsSummary requests a event tips summary.
// Request ID: k_EMsgClientToGCRequestEventTipsSummary
// Response ID: k_EMsgClientToGCRequestEventTipsSummaryResponse
// Request type: CMsgEventTipsSummaryRequest
// Response type: CMsgEventTipsSummaryResponse
func (d *Dota2) RequestEventTipsSummary(
	ctx context.Context,
	eventID protocol.EEvent,
	accountID uint32,
) (*protocol.CMsgEventTipsSummaryResponse, error) {
	req := &protocol.CMsgEventTipsSummaryRequest{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgEventTipsSummaryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummary),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummaryResponse),
		resp,
	)
}

// RequestFriendsPlayedCustomGame requests a friends played custom game.
// Request ID: k_EMsgClientToGCFriendsPlayedCustomGameRequest
// Response ID: k_EMsgGCToClientFriendsPlayedCustomGameResponse
// Request type: CMsgClientToGCFriendsPlayedCustomGameRequest
// Response type: CMsgGCToClientFriendsPlayedCustomGameResponse
func (d *Dota2) RequestFriendsPlayedCustomGame(
	ctx context.Context,
	customGameID uint64,
) (*protocol.CMsgGCToClientFriendsPlayedCustomGameResponse, error) {
	req := &protocol.CMsgClientToGCFriendsPlayedCustomGameRequest{
		CustomGameId: &customGameID,
	}
	resp := &protocol.CMsgGCToClientFriendsPlayedCustomGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFriendsPlayedCustomGameRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFriendsPlayedCustomGameResponse),
		resp,
	)
}

// RequestGetPlayerCardRoster requests to check if the target get player card roster.
// Request ID: k_EMsgClientToGCGetPlayerCardRosterRequest
// Response ID: k_EMsgClientToGCGetPlayerCardRosterResponse
// Request type: CMsgClientToGCGetPlayerCardRosterRequest
// Response type: CMsgClientToGCGetPlayerCardRosterResponse
func (d *Dota2) RequestGetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	timestamp uint32,
) (*protocol.CMsgClientToGCGetPlayerCardRosterResponse, error) {
	req := &protocol.CMsgClientToGCGetPlayerCardRosterRequest{
		LeagueId:  &leagueID,
		Timestamp: &timestamp,
	}
	resp := &protocol.CMsgClientToGCGetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterResponse),
		resp,
	)
}

// RequestGetRecentPlayTimeFriends requests to check if the target get recent play time friends.
// Request ID: k_EMsgGetRecentPlayTimeFriendsRequest
// Response ID: k_EMsgGetRecentPlayTimeFriendsResponse
// Request type: CMsgDOTAGetRecentPlayTimeFriendsRequest
// Response type: CMsgDOTAGetRecentPlayTimeFriendsResponse
func (d *Dota2) RequestGetRecentPlayTimeFriends(
	ctx context.Context,
) (*protocol.CMsgDOTAGetRecentPlayTimeFriendsResponse, error) {
	req := &protocol.CMsgDOTAGetRecentPlayTimeFriendsRequest{}
	resp := &protocol.CMsgDOTAGetRecentPlayTimeFriendsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGetRecentPlayTimeFriendsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGetRecentPlayTimeFriendsResponse),
		resp,
	)
}

// RequestGuildData requests a guild data.
// Request ID: k_EMsgClientToGCRequestGuildData
// Response ID: k_EMsgClientToGCRequestGuildDataResponse
// Request type: CMsgClientToGCRequestGuildData
// Response type: CMsgClientToGCRequestGuildDataResponse
func (d *Dota2) RequestGuildData(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCRequestGuildDataResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildData{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildDataResponse),
		resp,
	)
}

// RequestGuildEventMembers requests guild event members.
// Request ID: k_EMsgClientToGCRequestGuildEventMembers
// Response ID: k_EMsgClientToGCRequestGuildEventMembersResponse
// Request type: CMsgClientToGCRequestGuildEventMembers
// Response type: CMsgClientToGCRequestGuildEventMembersResponse
func (d *Dota2) RequestGuildEventMembers(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestGuildEventMembersResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildEventMembers{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildEventMembersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildEventMembers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildEventMembersResponse),
		resp,
	)
}

// RequestGuildFeed requests a guild feed.
// Request ID: k_EMsgClientToGCRequestGuildFeed
// Response ID: k_EMsgClientToGCRequestGuildFeedResponse
// Request type: CMsgClientToGCGuildFeedRequest
// Response type: CMsgClientToGCRequestGuildFeedResponse
func (d *Dota2) RequestGuildFeed(
	ctx context.Context,
	guildID uint32,
	lastSeenID uint64,
) (*protocol.CMsgClientToGCRequestGuildFeedResponse, error) {
	req := &protocol.CMsgClientToGCGuildFeedRequest{
		GuildId:    &guildID,
		LastSeenId: &lastSeenID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildFeedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeed),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeedResponse),
		resp,
	)
}

// RequestGuildMembership requests a guild membership.
// Request ID: k_EMsgClientToGCRequestGuildMembership
// Response ID: k_EMsgClientToGCRequestGuildMembershipResponse
// Request type: CMsgClientToGCRequestGuildMembership
// Response type: CMsgClientToGCRequestGuildMembershipResponse
func (d *Dota2) RequestGuildMembership(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestGuildMembershipResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildMembership{}
	resp := &protocol.CMsgClientToGCRequestGuildMembershipResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildMembership),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildMembershipResponse),
		resp,
	)
}

// RequestHeroGlobalData requests a hero global data.
// Request ID: k_EMsgHeroGlobalDataRequest
// Response ID: k_EMsgHeroGlobalDataResponse
// Request type: CMsgHeroGlobalDataRequest
// Response type: CMsgHeroGlobalDataResponse
func (d *Dota2) RequestHeroGlobalData(
	ctx context.Context,
	heroID uint32,
) (*protocol.CMsgHeroGlobalDataResponse, error) {
	req := &protocol.CMsgHeroGlobalDataRequest{
		HeroId: &heroID,
	}
	resp := &protocol.CMsgHeroGlobalDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgHeroGlobalDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgHeroGlobalDataResponse),
		resp,
	)
}

// RequestItemEditorReservations requests item editor reservations.
// Request ID: k_EMsgGCItemEditorReservationsRequest
// Response ID: k_EMsgGCItemEditorReservationsResponse
// Request type: CMsgGCItemEditorReservationsRequest
// Response type: CMsgGCItemEditorReservationsResponse
func (d *Dota2) RequestItemEditorReservations(
	ctx context.Context,
) (*protocol.CMsgGCItemEditorReservationsResponse, error) {
	req := &protocol.CMsgGCItemEditorReservationsRequest{}
	resp := &protocol.CMsgGCItemEditorReservationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReservationsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReservationsResponse),
		resp,
	)
}

// RequestJoinableCustomGameModes requests joinable custom game modes.
// Request ID: k_EMsgGCJoinableCustomGameModesRequest
// Response ID: k_EMsgGCJoinableCustomGameModesResponse
// Request type: CMsgJoinableCustomGameModesRequest
// Response type: CMsgJoinableCustomGameModesResponse
func (d *Dota2) RequestJoinableCustomGameModes(
	ctx context.Context,
	serverRegion uint32,
) (*protocol.CMsgJoinableCustomGameModesResponse, error) {
	req := &protocol.CMsgJoinableCustomGameModesRequest{
		ServerRegion: &serverRegion,
	}
	resp := &protocol.CMsgJoinableCustomGameModesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesResponse),
		resp,
	)
}

// RequestJoinableCustomLobbies requests joinable custom lobbies.
// Request ID: k_EMsgGCJoinableCustomLobbiesRequest
// Response ID: k_EMsgGCJoinableCustomLobbiesResponse
// Request type: CMsgJoinableCustomLobbiesRequest
// Response type: CMsgJoinableCustomLobbiesResponse
func (d *Dota2) RequestJoinableCustomLobbies(
	ctx context.Context,
	serverRegion uint32,
	customGameID uint64,
) (*protocol.CMsgJoinableCustomLobbiesResponse, error) {
	req := &protocol.CMsgJoinableCustomLobbiesRequest{
		ServerRegion: &serverRegion,
		CustomGameId: &customGameID,
	}
	resp := &protocol.CMsgJoinableCustomLobbiesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesResponse),
		resp,
	)
}

// RequestLatestConductScorecard requests a latest conduct scorecard.
// Request ID: k_EMsgClientToGCLatestConductScorecardRequest
// Response ID: k_EMsgClientToGCLatestConductScorecard
// Request type: CMsgPlayerConductScorecardRequest
// Response type: CMsgPlayerConductScorecard
func (d *Dota2) RequestLatestConductScorecard(
	ctx context.Context,
) (*protocol.CMsgPlayerConductScorecard, error) {
	req := &protocol.CMsgPlayerConductScorecardRequest{}
	resp := &protocol.CMsgPlayerConductScorecard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard),
		resp,
	)
}

// RequestLeagueAvailableLobbyNodes requests league available lobby nodes.
// Request ID: k_EMsgDOTALeagueAvailableLobbyNodesRequest
// Response ID: k_EMsgDOTALeagueAvailableLobbyNodes
// Request type: CMsgDOTALeagueAvailableLobbyNodesRequest
// Response type: CMsgDOTALeagueAvailableLobbyNodes
func (d *Dota2) RequestLeagueAvailableLobbyNodes(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgDOTALeagueAvailableLobbyNodes, error) {
	req := &protocol.CMsgDOTALeagueAvailableLobbyNodesRequest{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgDOTALeagueAvailableLobbyNodes{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodes),
		resp,
	)
}

// RequestMatchDetails requests match details.
// Request ID: k_EMsgGCMatchDetailsRequest
// Response ID: k_EMsgGCMatchDetailsResponse
// Request type: CMsgGCMatchDetailsRequest
// Response type: CMsgGCMatchDetailsResponse
func (d *Dota2) RequestMatchDetails(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgGCMatchDetailsResponse, error) {
	req := &protocol.CMsgGCMatchDetailsRequest{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgGCMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchDetailsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchDetailsResponse),
		resp,
	)
}

// RequestMatchesMinimal requests a matches minimal.
// Request ID: k_EMsgClientToGCMatchesMinimalRequest
// Response ID: k_EMsgClientToGCMatchesMinimalResponse
// Request type: CMsgClientToGCMatchesMinimalRequest
// Response type: CMsgClientToGCMatchesMinimalResponse
func (d *Dota2) RequestMatchesMinimal(
	ctx context.Context,
	matchIDs []uint64,
) (*protocol.CMsgClientToGCMatchesMinimalResponse, error) {
	req := &protocol.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}
	resp := &protocol.CMsgClientToGCMatchesMinimalResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse),
		resp,
	)
}

// RequestMatchmakingStats requests matchmaking stats.
// Request ID: k_EMsgGCMatchmakingStatsRequest
// Response ID: k_EMsgGCMatchmakingStatsResponse
// Request type: CMsgDOTAMatchmakingStatsRequest
// Response type: CMsgDOTAMatchmakingStatsResponse
func (d *Dota2) RequestMatchmakingStats(
	ctx context.Context,
) (*protocol.CMsgDOTAMatchmakingStatsResponse, error) {
	req := &protocol.CMsgDOTAMatchmakingStatsRequest{}
	resp := &protocol.CMsgDOTAMatchmakingStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchmakingStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchmakingStatsResponse),
		resp,
	)
}

// RequestMyTeamInfo requests a my team info.
// Request ID: k_EMsgClientToGCMyTeamInfoRequest
// Response ID: k_EMsgGCToClientTeamInfo
// Request type: CMsgDOTAMyTeamInfoRequest
// Response type: CMsgDOTATeamInfo
func (d *Dota2) RequestMyTeamInfo(
	ctx context.Context,
) (*protocol.CMsgDOTATeamInfo, error) {
	req := &protocol.CMsgDOTAMyTeamInfoRequest{}
	resp := &protocol.CMsgDOTATeamInfo{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamInfo),
		resp,
	)
}

// RequestNotifications requests notifications.
// Request ID: k_EMsgGCNotificationsRequest
// Response ID: k_EMsgGCNotificationsResponse
// Request type: CMsgGCNotificationsRequest
// Response type: CMsgGCNotificationsResponse
func (d *Dota2) RequestNotifications(
	ctx context.Context,
) (*protocol.CMsgGCNotificationsResponse, error) {
	req := &protocol.CMsgGCNotificationsRequest{}
	resp := &protocol.CMsgGCNotificationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsResponse),
		resp,
	)
}

// RequestNotificationsMarkRead requests a notifications mark read.
// Request ID: k_EMsgGCNotificationsMarkReadRequest
// Request type: CMsgGCNotificationsMarkReadRequest
func (d *Dota2) RequestNotificationsMarkRead() {
	req := &protocol.CMsgGCNotificationsMarkReadRequest{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest), req)
}

// RequestPlayerCoachMatch requests a player coach match.
// Request ID: k_EMsgClientToGCRequestPlayerCoachMatch
// Response ID: k_EMsgClientToGCRequestPlayerCoachMatchResponse
// Request type: CMsgClientToGCRequestPlayerCoachMatch
// Response type: CMsgClientToGCRequestPlayerCoachMatchResponse
func (d *Dota2) RequestPlayerCoachMatch(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCRequestPlayerCoachMatchResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerCoachMatch{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerCoachMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatchResponse),
		resp,
	)
}

// RequestPlayerCoachMatches requests player coach matches.
// Request ID: k_EMsgClientToGCRequestPlayerCoachMatches
// Response ID: k_EMsgClientToGCRequestPlayerCoachMatchesResponse
// Request type: CMsgClientToGCRequestPlayerCoachMatches
// Response type: CMsgClientToGCRequestPlayerCoachMatchesResponse
func (d *Dota2) RequestPlayerCoachMatches(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestPlayerCoachMatchesResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerCoachMatches{}
	resp := &protocol.CMsgClientToGCRequestPlayerCoachMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatches),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatchesResponse),
		resp,
	)
}

// RequestPlayerHeroRecentAccomplishments requests player hero recent accomplishments.
// Request ID: k_EMsgClientToGCRequestPlayerHeroRecentAccomplishments
// Response ID: k_EMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse
// Request type: CMsgClientToGCRequestPlayerHeroRecentAccomplishments
// Response type: CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse
func (d *Dota2) RequestPlayerHeroRecentAccomplishments(
	ctx context.Context,
	accountID uint32,
	heroID uint32,
) (*protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishments{
		AccountId: &accountID,
		HeroId:    &heroID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerHeroRecentAccomplishments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse),
		resp,
	)
}

// RequestPlayerRecentAccomplishments requests player recent accomplishments.
// Request ID: k_EMsgClientToGCRequestPlayerRecentAccomplishments
// Response ID: k_EMsgClientToGCRequestPlayerRecentAccomplishmentsResponse
// Request type: CMsgClientToGCRequestPlayerRecentAccomplishments
// Response type: CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse
func (d *Dota2) RequestPlayerRecentAccomplishments(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerRecentAccomplishments{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerRecentAccomplishments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerRecentAccomplishmentsResponse),
		resp,
	)
}

// RequestPlayerStats requests player stats.
// Request ID: k_EMsgClientToGCPlayerStatsRequest
// Response ID: k_EMsgGCToClientPlayerStatsResponse
// Request type: CMsgClientToGCPlayerStatsRequest
// Response type: CMsgGCToClientPlayerStatsResponse
func (d *Dota2) RequestPlayerStats(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgGCToClientPlayerStatsResponse, error) {
	req := &protocol.CMsgClientToGCPlayerStatsRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgGCToClientPlayerStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlayerStatsResponse),
		resp,
	)
}

// RequestPlusWeeklyChallengeResult requests a plus weekly challenge result.
// Request ID: k_EMsgClientToGCRequestPlusWeeklyChallengeResult
// Response ID: k_EMsgClientToGCRequestPlusWeeklyChallengeResultResponse
// Request type: CMsgClientToGCRequestPlusWeeklyChallengeResult
// Response type: CMsgClientToGCRequestPlusWeeklyChallengeResultResponse
func (d *Dota2) RequestPlusWeeklyChallengeResult(
	ctx context.Context,
	eventID protocol.EEvent,
	week uint32,
) (*protocol.CMsgClientToGCRequestPlusWeeklyChallengeResultResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlusWeeklyChallengeResult{
		EventId: &eventID,
		Week:    &week,
	}
	resp := &protocol.CMsgClientToGCRequestPlusWeeklyChallengeResultResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlusWeeklyChallengeResult),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlusWeeklyChallengeResultResponse),
		resp,
	)
}

// RequestPrivateCoachingSession requests a private coaching session.
// Request ID: k_EMsgClientToGCRequestPrivateCoachingSession
// Response ID: k_EMsgClientToGCRequestPrivateCoachingSessionResponse
// Request type: CMsgClientToGCRequestPrivateCoachingSession
// Response type: CMsgClientToGCRequestPrivateCoachingSessionResponse
func (d *Dota2) RequestPrivateCoachingSession(
	ctx context.Context,
	language uint32,
) (*protocol.CMsgClientToGCRequestPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCRequestPrivateCoachingSession{
		Language: &language,
	}
	resp := &protocol.CMsgClientToGCRequestPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPrivateCoachingSessionResponse),
		resp,
	)
}

// RequestPrivateMetadataKey requests a private metadata key.
// Request ID: k_EMsgPrivateMetadataKeyRequest
// Response ID: k_EMsgPrivateMetadataKeyResponse
// Request type: CMsgPrivateMetadataKeyRequest
// Response type: CMsgPrivateMetadataKeyResponse
func (d *Dota2) RequestPrivateMetadataKey(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgPrivateMetadataKeyResponse, error) {
	req := &protocol.CMsgPrivateMetadataKeyRequest{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgPrivateMetadataKeyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPrivateMetadataKeyRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPrivateMetadataKeyResponse),
		resp,
	)
}

// RequestProfile requests a profile.
// Request ID: k_EMsgProfileRequest
// Response ID: k_EMsgProfileResponse
// Request type: CMsgProfileRequest
// Response type: CMsgProfileResponse
func (d *Dota2) RequestProfile(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgProfileResponse, error) {
	req := &protocol.CMsgProfileRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgProfileResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileResponse),
		resp,
	)
}

// RequestQuickStats requests quick stats.
// Request ID: k_EMsgClientToGCQuickStatsRequest
// Response ID: k_EMsgClientToGCQuickStatsResponse
// Request type: CMsgDOTAClientToGCQuickStatsRequest
// Response type: CMsgDOTAClientToGCQuickStatsResponse
func (d *Dota2) RequestQuickStats(
	ctx context.Context,
	playerAccountID uint32,
	heroID uint32,
	itemID uint32,
	leagueID uint32,
) (*protocol.CMsgDOTAClientToGCQuickStatsResponse, error) {
	req := &protocol.CMsgDOTAClientToGCQuickStatsRequest{
		PlayerAccountId: &playerAccountID,
		HeroId:          &heroID,
		ItemId:          &itemID,
		LeagueId:        &leagueID,
	}
	resp := &protocol.CMsgDOTAClientToGCQuickStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCQuickStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCQuickStatsResponse),
		resp,
	)
}

// RequestReportsRemaining requests a reports remaining.
// Request ID: k_EMsgGCReportsRemainingRequest
// Response ID: k_EMsgGCReportsRemainingResponse
// Request type: CMsgDOTAReportsRemainingRequest
// Response type: CMsgDOTAReportsRemainingResponse
func (d *Dota2) RequestReportsRemaining(
	ctx context.Context,
) (*protocol.CMsgDOTAReportsRemainingResponse, error) {
	req := &protocol.CMsgDOTAReportsRemainingRequest{}
	resp := &protocol.CMsgDOTAReportsRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCReportsRemainingRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCReportsRemainingResponse),
		resp,
	)
}

// RequestRespondToCoachFriend requests a respond to coach friend.
// Request ID: k_EMsgClientToGCRespondToCoachFriendRequest
// Response ID: k_EMsgClientToGCRespondToCoachFriendRequestResponse
// Request type: CMsgClientToGCRespondToCoachFriendRequest
// Response type: CMsgClientToGCRespondToCoachFriendRequestResponse
func (d *Dota2) RequestRespondToCoachFriend(
	ctx context.Context,
	coachAccountID uint32,
	response protocol.ELobbyMemberCoachRequestState,
) (*protocol.CMsgClientToGCRespondToCoachFriendRequestResponse, error) {
	req := &protocol.CMsgClientToGCRespondToCoachFriendRequest{
		CoachAccountId: &coachAccountID,
		Response:       &response,
	}
	resp := &protocol.CMsgClientToGCRespondToCoachFriendRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRespondToCoachFriendRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRespondToCoachFriendRequestResponse),
		resp,
	)
}

// RequestSelectionPriorityChoice requests a selection priority choice.
// Request ID: k_EMsgSelectionPriorityChoiceRequest
// Response ID: k_EMsgSelectionPriorityChoiceResponse
// Request type: CMsgDOTASelectionPriorityChoiceRequest
// Response type: CMsgDOTASelectionPriorityChoiceResponse
func (d *Dota2) RequestSelectionPriorityChoice(
	ctx context.Context,
	choice protocol.DOTASelectionPriorityChoice,
) (*protocol.CMsgDOTASelectionPriorityChoiceResponse, error) {
	req := &protocol.CMsgDOTASelectionPriorityChoiceRequest{
		Choice: &choice,
	}
	resp := &protocol.CMsgDOTASelectionPriorityChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceResponse),
		resp,
	)
}

// RequestSetPlayerCardRoster requests to check if the target set player card roster.
// Request ID: k_EMsgClientToGCSetPlayerCardRosterRequest
// Response ID: k_EMsgClientToGCSetPlayerCardRosterResponse
// Request type: CMsgClientToGCSetPlayerCardRosterRequest
// Response type: CMsgClientToGCSetPlayerCardRosterResponse
func (d *Dota2) RequestSetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	timestamp uint32,
	slot uint32,
	playerCardItemID uint64,
	eventID uint32,
) (*protocol.CMsgClientToGCSetPlayerCardRosterResponse, error) {
	req := &protocol.CMsgClientToGCSetPlayerCardRosterRequest{
		LeagueId:         &leagueID,
		Timestamp:        &timestamp,
		Slot:             &slot,
		PlayerCardItemId: &playerCardItemID,
		EventId:          &eventID,
	}
	resp := &protocol.CMsgClientToGCSetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterResponse),
		resp,
	)
}

// RequestSocialFeed requests a social feed.
// Request ID: k_EMsgClientToGCRequestSocialFeed
// Response ID: k_EMsgClientToGCRequestSocialFeedResponse
// Request type: CMsgSocialFeedRequest
// Response type: CMsgSocialFeedResponse
func (d *Dota2) RequestSocialFeed(
	ctx context.Context,
	accountID uint32,
	selfOnly bool,
) (*protocol.CMsgSocialFeedResponse, error) {
	req := &protocol.CMsgSocialFeedRequest{
		AccountId: &accountID,
		SelfOnly:  &selfOnly,
	}
	resp := &protocol.CMsgSocialFeedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeed),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedResponse),
		resp,
	)
}

// RequestSocialFeedComments requests social feed comments.
// Request ID: k_EMsgClientToGCRequestSocialFeedComments
// Response ID: k_EMsgClientToGCRequestSocialFeedCommentsResponse
// Request type: CMsgSocialFeedCommentsRequest
// Response type: CMsgSocialFeedCommentsResponse
func (d *Dota2) RequestSocialFeedComments(
	ctx context.Context,
	feedEventID uint64,
) (*protocol.CMsgSocialFeedCommentsResponse, error) {
	req := &protocol.CMsgSocialFeedCommentsRequest{
		FeedEventId: &feedEventID,
	}
	resp := &protocol.CMsgSocialFeedCommentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedComments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedCommentsResponse),
		resp,
	)
}

// RequestSocialFeedPostComment requests a social feed post comment.
// Request ID: k_EMsgClientToGCSocialFeedPostCommentRequest
// Response ID: k_EMsgGCToClientSocialFeedPostCommentResponse
// Request type: CMsgClientToGCSocialFeedPostCommentRequest
// Response type: CMsgGCToClientSocialFeedPostCommentResponse
func (d *Dota2) RequestSocialFeedPostComment(
	ctx context.Context,
	eventID uint64,
	comment string,
) (*protocol.CMsgGCToClientSocialFeedPostCommentResponse, error) {
	req := &protocol.CMsgClientToGCSocialFeedPostCommentRequest{
		EventId: &eventID,
		Comment: &comment,
	}
	resp := &protocol.CMsgGCToClientSocialFeedPostCommentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostCommentRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostCommentResponse),
		resp,
	)
}

// RequestSocialFeedPostMessage requests a social feed post message.
// Request ID: k_EMsgClientToGCSocialFeedPostMessageRequest
// Response ID: k_EMsgGCToClientSocialFeedPostMessageResponse
// Request type: CMsgClientToGCSocialFeedPostMessageRequest
// Response type: CMsgGCToClientSocialFeedPostMessageResponse
func (d *Dota2) RequestSocialFeedPostMessage(
	ctx context.Context,
	message string,
	matchID uint64,
	matchTimestamp uint32,
) (*protocol.CMsgGCToClientSocialFeedPostMessageResponse, error) {
	req := &protocol.CMsgClientToGCSocialFeedPostMessageRequest{
		Message:        &message,
		MatchId:        &matchID,
		MatchTimestamp: &matchTimestamp,
	}
	resp := &protocol.CMsgGCToClientSocialFeedPostMessageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostMessageRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostMessageResponse),
		resp,
	)
}

// RequestSteamDatagramTicket requests a steam datagram ticket.
// Request ID: k_EMsgClientToGCRequestSteamDatagramTicket
// Response ID: k_EMsgClientToGCRequestSteamDatagramTicketResponse
// Request type: CMsgClientToGCRequestSteamDatagramTicket
// Response type: CMsgClientToGCRequestSteamDatagramTicketResponse
func (d *Dota2) RequestSteamDatagramTicket(
	ctx context.Context,
	serverSteamID steamid.SteamId,
) (*protocol.CMsgClientToGCRequestSteamDatagramTicketResponse, error) {
	serverSteamIDU64Val := uint64(serverSteamID)
	serverSteamIDU64 := &serverSteamIDU64Val
	req := &protocol.CMsgClientToGCRequestSteamDatagramTicket{
		ServerSteamId: serverSteamIDU64,
	}
	resp := &protocol.CMsgClientToGCRequestSteamDatagramTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicket),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicketResponse),
		resp,
	)
}

// RequestSubmitPlayerAvoid requests to check if the target submit player avoid.
// Request ID: k_EMsgGCSubmitPlayerAvoidRequest
// Response ID: k_EMsgGCSubmitPlayerAvoidRequestResponse
// Request type: CMsgDOTASubmitPlayerAvoidRequest
// Response type: CMsgDOTASubmitPlayerAvoidRequestResponse
func (d *Dota2) RequestSubmitPlayerAvoid(
	ctx context.Context,
	targetAccountID uint32,
	lobbyID uint64,
) (*protocol.CMsgDOTASubmitPlayerAvoidRequestResponse, error) {
	req := &protocol.CMsgDOTASubmitPlayerAvoidRequest{
		TargetAccountId: &targetAccountID,
		LobbyId:         &lobbyID,
	}
	resp := &protocol.CMsgDOTASubmitPlayerAvoidRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerAvoidRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerAvoidRequestResponse),
		resp,
	)
}

// RequestTeammateStats requests teammate stats.
// Request ID: k_EMsgClientToGCTeammateStatsRequest
// Response ID: k_EMsgClientToGCTeammateStatsResponse
// Request type: CMsgClientToGCTeammateStatsRequest
// Response type: CMsgClientToGCTeammateStatsResponse
func (d *Dota2) RequestTeammateStats(
	ctx context.Context,
) (*protocol.CMsgClientToGCTeammateStatsResponse, error) {
	req := &protocol.CMsgClientToGCTeammateStatsRequest{}
	resp := &protocol.CMsgClientToGCTeammateStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsResponse),
		resp,
	)
}

// RequestTopFriendMatches requests top friend matches.
// Request ID: k_EMsgClientToGCTopFriendMatchesRequest
// Response ID: k_EMsgGCToClientTopFriendMatchesResponse
// Request type: CMsgClientToGCTopFriendMatchesRequest
// Response type: CMsgGCToClientTopFriendMatchesResponse
func (d *Dota2) RequestTopFriendMatches(
	ctx context.Context,
) (*protocol.CMsgGCToClientTopFriendMatchesResponse, error) {
	req := &protocol.CMsgClientToGCTopFriendMatchesRequest{}
	resp := &protocol.CMsgGCToClientTopFriendMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTopFriendMatchesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTopFriendMatchesResponse),
		resp,
	)
}

// RequestTopLeagueMatches requests top league matches.
// Request ID: k_EMsgClientToGCTopLeagueMatchesRequest
// Response ID: k_EMsgGCToClientTopLeagueMatchesResponse
// Request type: CMsgClientToGCTopLeagueMatchesRequest
// Response type: CMsgGCToClientTopLeagueMatchesResponse
func (d *Dota2) RequestTopLeagueMatches(
	ctx context.Context,
) (*protocol.CMsgGCToClientTopLeagueMatchesResponse, error) {
	req := &protocol.CMsgClientToGCTopLeagueMatchesRequest{}
	resp := &protocol.CMsgGCToClientTopLeagueMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTopLeagueMatchesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTopLeagueMatchesResponse),
		resp,
	)
}

// RequestTransferSeasonalMMR requests to check if the target transfer seasonal mmr.
// Request ID: k_EMsgClientToGCTransferSeasonalMMRRequest
// Response ID: k_EMsgClientToGCTransferSeasonalMMRResponse
// Request type: CMsgClientToGCTransferSeasonalMMRRequest
// Response type: CMsgClientToGCTransferSeasonalMMRResponse
func (d *Dota2) RequestTransferSeasonalMMR(
	ctx context.Context,
	isParty bool,
) (*protocol.CMsgClientToGCTransferSeasonalMMRResponse, error) {
	req := &protocol.CMsgClientToGCTransferSeasonalMMRRequest{
		IsParty: &isParty,
	}
	resp := &protocol.CMsgClientToGCTransferSeasonalMMRResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRResponse),
		resp,
	)
}

// RequestUnanchorPhoneNumber requests a unanchor phone number.
// Request ID: k_EMsgUnanchorPhoneNumberRequest
// Response ID: k_EMsgUnanchorPhoneNumberResponse
// Request type: CMsgDOTAUnanchorPhoneNumberRequest
// Response type: CMsgDOTAUnanchorPhoneNumberResponse
func (d *Dota2) RequestUnanchorPhoneNumber(
	ctx context.Context,
) (*protocol.CMsgDOTAUnanchorPhoneNumberResponse, error) {
	req := &protocol.CMsgDOTAUnanchorPhoneNumberRequest{}
	resp := &protocol.CMsgDOTAUnanchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberResponse),
		resp,
	)
}

// RequestUnderDraft requests a under draft.
// Request ID: k_EMsgClientToGCUnderDraftRequest
// Response ID: k_EMsgClientToGCUnderDraftResponse
// Request type: CMsgClientToGCUnderDraftRequest
// Response type: CMsgClientToGCUnderDraftResponse
func (d *Dota2) RequestUnderDraft(
	ctx context.Context,
	accountID uint32,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRequest{
		AccountId: &accountID,
		EventId:   &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftResponse),
		resp,
	)
}

// RequestWagering requests a wagering.
// Request ID: k_EMsgClientToGCWageringRequest
// Response ID: k_EMsgGCToClientWageringResponse
// Request type: CMsgClientToGCWageringRequest
// Response type: CMsgGCToClientWageringResponse
func (d *Dota2) RequestWagering(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgGCToClientWageringResponse, error) {
	req := &protocol.CMsgClientToGCWageringRequest{
		EventId: &eventID,
	}
	resp := &protocol.CMsgGCToClientWageringResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWageringRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientWageringResponse),
		resp,
	)
}

// RerollDraftUnder rerolls a draft under.
// Request ID: k_EMsgClientToGCUnderDraftReroll
// Response ID: k_EMsgClientToGCUnderDraftRerollResponse
// Request type: CMsgClientToGCUnderDraftReroll
// Response type: CMsgClientToGCUnderDraftRerollResponse
func (d *Dota2) RerollDraftUnder(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftRerollResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftReroll{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRerollResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftReroll),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRerollResponse),
		resp,
	)
}

// RerollPlayerChallenge rerolls a player challenge.
// Request ID: k_EMsgClientToGCRerollPlayerChallenge
// Request type: CMsgClientToGCRerollPlayerChallenge
func (d *Dota2) RerollPlayerChallenge(
	eventID protocol.EEvent,
	sequenceID uint32,
	heroID uint32,
) {
	req := &protocol.CMsgClientToGCRerollPlayerChallenge{
		EventId:    &eventID,
		SequenceId: &sequenceID,
		HeroId:     &heroID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRerollPlayerChallenge), req)
}

// ReserveEditorItemItemDef reserves a editor item item def.
// Request ID: k_EMsgGCItemEditorReserveItemDef
// Response ID: k_EMsgGCItemEditorReserveItemDefResponse
// Request type: CMsgGCItemEditorReserveItemDef
// Response type: CMsgGCItemEditorReserveItemDefResponse
func (d *Dota2) ReserveEditorItemItemDef(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*protocol.CMsgGCItemEditorReserveItemDefResponse, error) {
	req := &protocol.CMsgGCItemEditorReserveItemDef{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &protocol.CMsgGCItemEditorReserveItemDefResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDef),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDefResponse),
		resp,
	)
}

// RespondToTeamInvite is undocumented.
// Request ID: k_EMsgGCTeamInvite_InviteeResponseToGC
// Request type: CMsgDOTATeamInvite_InviteeResponseToGC
func (d *Dota2) RespondToTeamInvite(
	result protocol.ETeamInviteResult,
) {
	req := &protocol.CMsgDOTATeamInvite_InviteeResponseToGC{
		Result: &result,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC), req)
}

// SelectCompendiumInGamePrediction selects a compendium in game prediction.
// Request ID: k_EMsgClientToGCSelectCompendiumInGamePrediction
// Response ID: k_EMsgClientToGCSelectCompendiumInGamePredictionResponse
// Request type: CMsgClientToGCSelectCompendiumInGamePrediction
// Response type: CMsgClientToGCSelectCompendiumInGamePredictionResponse
func (d *Dota2) SelectCompendiumInGamePrediction(
	ctx context.Context,
	matchID uint64,
	predictions []*protocol.CMsgClientToGCSelectCompendiumInGamePrediction_Prediction,
	leagueID uint32,
) (*protocol.CMsgClientToGCSelectCompendiumInGamePredictionResponse, error) {
	req := &protocol.CMsgClientToGCSelectCompendiumInGamePrediction{
		MatchId:     &matchID,
		Predictions: predictions,
		LeagueId:    &leagueID,
	}
	resp := &protocol.CMsgClientToGCSelectCompendiumInGamePredictionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePrediction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePredictionResponse),
		resp,
	)
}

// SelectGuildContract selects a guild contract.
// Request ID: k_EMsgClientToGCSelectGuildContract
// Response ID: k_EMsgClientToGCSelectGuildContractResponse
// Request type: CMsgClientToGCSelectGuildContract
// Response type: CMsgClientToGCSelectGuildContractResponse
func (d *Dota2) SelectGuildContract(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
	contractID uint64,
	contractSlot uint32,
) (*protocol.CMsgClientToGCSelectGuildContractResponse, error) {
	req := &protocol.CMsgClientToGCSelectGuildContract{
		GuildId:      &guildID,
		EventId:      &eventID,
		ContractId:   &contractID,
		ContractSlot: &contractSlot,
	}
	resp := &protocol.CMsgClientToGCSelectGuildContractResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectGuildContract),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectGuildContractResponse),
		resp,
	)
}

// SendAcceptInviteToGuild sends a accept invite to guild.
// Request ID: k_EMsgClientToGCAcceptInviteToGuild
// Response ID: k_EMsgClientToGCAcceptInviteToGuildResponse
// Request type: CMsgClientToGCAcceptInviteToGuild
// Response type: CMsgClientToGCAcceptInviteToGuildResponse
func (d *Dota2) SendAcceptInviteToGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCAcceptInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCAcceptInviteToGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCAcceptInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptInviteToGuildResponse),
		resp,
	)
}

// SendAcceptPrivateCoachingSession sends a accept private coaching session.
// Request ID: k_EMsgClientToGCAcceptPrivateCoachingSession
// Response ID: k_EMsgClientToGCAcceptPrivateCoachingSessionResponse
// Request type: CMsgClientToGCAcceptPrivateCoachingSession
// Response type: CMsgClientToGCAcceptPrivateCoachingSessionResponse
func (d *Dota2) SendAcceptPrivateCoachingSession(
	ctx context.Context,
	coachingSessionID uint64,
) (*protocol.CMsgClientToGCAcceptPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCAcceptPrivateCoachingSession{
		CoachingSessionId: &coachingSessionID,
	}
	resp := &protocol.CMsgClientToGCAcceptPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptPrivateCoachingSessionResponse),
		resp,
	)
}

// SendAddGuildRole sends a add guild role.
// Request ID: k_EMsgClientToGCAddGuildRole
// Response ID: k_EMsgClientToGCAddGuildRoleResponse
// Request type: CMsgClientToGCAddGuildRole
// Response type: CMsgClientToGCAddGuildRoleResponse
func (d *Dota2) SendAddGuildRole(
	ctx context.Context,
	guildID uint32,
	roleName string,
	roleFlags uint32,
) (*protocol.CMsgClientToGCAddGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCAddGuildRole{
		GuildId:   &guildID,
		RoleName:  &roleName,
		RoleFlags: &roleFlags,
	}
	resp := &protocol.CMsgClientToGCAddGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddGuildRoleResponse),
		resp,
	)
}

// SendAddPlayerToGuildChat sends a add player to guild chat.
// Request ID: k_EMsgClientToGCAddPlayerToGuildChat
// Response ID: k_EMsgClientToGCAddPlayerToGuildChatResponse
// Request type: CMsgClientToGCAddPlayerToGuildChat
// Response type: CMsgClientToGCAddPlayerToGuildChatResponse
func (d *Dota2) SendAddPlayerToGuildChat(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCAddPlayerToGuildChatResponse, error) {
	req := &protocol.CMsgClientToGCAddPlayerToGuildChat{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCAddPlayerToGuildChatResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddPlayerToGuildChat),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddPlayerToGuildChatResponse),
		resp,
	)
}

// SendBalancedShuffleLobby sends a balanced shuffle lobby.
// Request ID: k_EMsgGCBalancedShuffleLobby
// Request type: CMsgBalancedShuffleLobby
func (d *Dota2) SendBalancedShuffleLobby() {
	req := &protocol.CMsgBalancedShuffleLobby{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby), req)
}

// SendCavernCrawlClaimRoom sends a cavern crawl claim room.
// Request ID: k_EMsgClientToGCCavernCrawlClaimRoom
// Response ID: k_EMsgClientToGCCavernCrawlClaimRoomResponse
// Request type: CMsgClientToGCCavernCrawlClaimRoom
// Response type: CMsgClientToGCCavernCrawlClaimRoomResponse
func (d *Dota2) SendCavernCrawlClaimRoom(
	ctx context.Context,
	eventID uint32,
	roomID uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlClaimRoomResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlClaimRoom{
		EventId:    &eventID,
		RoomId:     &roomID,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlClaimRoomResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlClaimRoom),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlClaimRoomResponse),
		resp,
	)
}

// SendCavernCrawlUseItemOnPath sends a cavern crawl use item on path.
// Request ID: k_EMsgClientToGCCavernCrawlUseItemOnPath
// Response ID: k_EMsgClientToGCCavernCrawlUseItemOnPathResponse
// Request type: CMsgClientToGCCavernCrawlUseItemOnPath
// Response type: CMsgClientToGCCavernCrawlUseItemOnPathResponse
func (d *Dota2) SendCavernCrawlUseItemOnPath(
	ctx context.Context,
	eventID uint32,
	pathID uint32,
	itemType uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlUseItemOnPathResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlUseItemOnPath{
		EventId:    &eventID,
		PathId:     &pathID,
		ItemType:   &itemType,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlUseItemOnPathResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnPath),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnPathResponse),
		resp,
	)
}

// SendCavernCrawlUseItemOnRoom sends a cavern crawl use item on room.
// Request ID: k_EMsgClientToGCCavernCrawlUseItemOnRoom
// Response ID: k_EMsgClientToGCCavernCrawlUseItemOnRoomResponse
// Request type: CMsgClientToGCCavernCrawlUseItemOnRoom
// Response type: CMsgClientToGCCavernCrawlUseItemOnRoomResponse
func (d *Dota2) SendCavernCrawlUseItemOnRoom(
	ctx context.Context,
	eventID uint32,
	roomID uint32,
	itemType uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlUseItemOnRoomResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlUseItemOnRoom{
		EventId:    &eventID,
		RoomId:     &roomID,
		ItemType:   &itemType,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlUseItemOnRoomResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnRoom),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnRoomResponse),
		resp,
	)
}

// SendChatMessage sends a chat message.
// Request ID: k_EMsgGCChatMessage
// Request type: CMsgDOTAChatMessage
func (d *Dota2) SendChatMessage(
	req *protocol.CMsgDOTAChatMessage,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCChatMessage), req)
}

// SendClaimEventAction sends a claim event action.
// Request ID: k_EMsgDOTAClaimEventAction
// Response ID: k_EMsgDOTAClaimEventActionResponse
// Request type: CMsgDOTAClaimEventAction
// Response type: CMsgDOTAClaimEventActionResponse
func (d *Dota2) SendClaimEventAction(
	ctx context.Context,
	req *protocol.CMsgDOTAClaimEventAction,
) (*protocol.CMsgDOTAClaimEventActionResponse, error) {
	resp := &protocol.CMsgDOTAClaimEventActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAClaimEventAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAClaimEventActionResponse),
		resp,
	)
}

// SendClaimEventActionUsingItem sends a claim event action using item.
// Request ID: k_EMsgClientToGCClaimEventActionUsingItem
// Response ID: k_EMsgClientToGCClaimEventActionUsingItemResponse
// Request type: CMsgClientToGCClaimEventActionUsingItem
// Response type: CMsgClientToGCClaimEventActionUsingItemResponse
func (d *Dota2) SendClaimEventActionUsingItem(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
	itemID uint64,
	quantity uint32,
) (*protocol.CMsgClientToGCClaimEventActionUsingItemResponse, error) {
	req := &protocol.CMsgClientToGCClaimEventActionUsingItem{
		EventId:  &eventID,
		ActionId: &actionID,
		ItemId:   &itemID,
		Quantity: &quantity,
	}
	resp := &protocol.CMsgClientToGCClaimEventActionUsingItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimEventActionUsingItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimEventActionUsingItemResponse),
		resp,
	)
}

// SendClaimLeaderboardRewards sends claim leaderboard rewards.
// Request ID: k_EMsgClientToGCClaimLeaderboardRewards
// Response ID: k_EMsgClientToGCClaimLeaderboardRewardsResponse
// Request type: CMsgClientToGCClaimLeaderboardRewards
// Response type: CMsgClientToGCClaimLeaderboardRewardsResponse
func (d *Dota2) SendClaimLeaderboardRewards(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCClaimLeaderboardRewardsResponse, error) {
	req := &protocol.CMsgClientToGCClaimLeaderboardRewards{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCClaimLeaderboardRewardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimLeaderboardRewards),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimLeaderboardRewardsResponse),
		resp,
	)
}

// SendCoachFriend sends a coach friend.
// Request ID: k_EMsgClientToGCCoachFriend
// Response ID: k_EMsgClientToGCCoachFriendResponse
// Request type: CMsgClientToGCCoachFriend
// Response type: CMsgClientToGCCoachFriendResponse
func (d *Dota2) SendCoachFriend(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCCoachFriendResponse, error) {
	req := &protocol.CMsgClientToGCCoachFriend{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCCoachFriendResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCoachFriend),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCoachFriendResponse),
		resp,
	)
}

// SendCustomGameClientFinishedLoading sends a custom game client finished loading.
// Request ID: k_EMsgCustomGameClientFinishedLoading
// Request type: CMsgDOTACustomGameClientFinishedLoading
func (d *Dota2) SendCustomGameClientFinishedLoading(
	req *protocol.CMsgDOTACustomGameClientFinishedLoading,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgCustomGameClientFinishedLoading), req)
}

// SendCustomGameListenServerStartedLoading sends a custom game listen server started loading.
// Request ID: k_EMsgCustomGameListenServerStartedLoading
// Request type: CMsgDOTACustomGameListenServerStartedLoading
func (d *Dota2) SendCustomGameListenServerStartedLoading(
	lobbyID uint64,
	customGameID uint64,
	lobbyMembers []uint32,
	startTime uint32,
) {
	req := &protocol.CMsgDOTACustomGameListenServerStartedLoading{
		LobbyId:      &lobbyID,
		CustomGameId: &customGameID,
		LobbyMembers: lobbyMembers,
		StartTime:    &startTime,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgCustomGameListenServerStartedLoading), req)
}

// SendDeclineInviteToGuild sends a decline invite to guild.
// Request ID: k_EMsgClientToGCDeclineInviteToGuild
// Response ID: k_EMsgClientToGCDeclineInviteToGuildResponse
// Request type: CMsgClientToGCDeclineInviteToGuild
// Response type: CMsgClientToGCDeclineInviteToGuildResponse
func (d *Dota2) SendDeclineInviteToGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCDeclineInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCDeclineInviteToGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCDeclineInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeclineInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeclineInviteToGuildResponse),
		resp,
	)
}

// SendDetailedGameStats sends detailed game stats.
// Request ID: k_EMsgDetailedGameStats
// Request type: CMsgDetailedGameStats
func (d *Dota2) SendDetailedGameStats(
	req *protocol.CMsgDetailedGameStats,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDetailedGameStats), req)
}

// SendDevDeleteEventActions sends dev delete event actions.
// Request ID: k_EMsgDevDeleteEventActions
// Response ID: k_EMsgDevDeleteEventActionsResponse
// Request type: CMsgDevDeleteEventActions
// Response type: CMsgDevDeleteEventActionsResponse
func (d *Dota2) SendDevDeleteEventActions(
	ctx context.Context,
	eventID protocol.EEvent,
	startActionID uint32,
	endActionID uint32,
	removeAudit bool,
) (*protocol.CMsgDevDeleteEventActionsResponse, error) {
	req := &protocol.CMsgDevDeleteEventActions{
		EventId:       &eventID,
		StartActionId: &startActionID,
		EndActionId:   &endActionID,
		RemoveAudit:   &removeAudit,
	}
	resp := &protocol.CMsgDevDeleteEventActionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevDeleteEventActions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevDeleteEventActionsResponse),
		resp,
	)
}

// SendDevResetEventState sends a dev reset event state.
// Request ID: k_EMsgDevResetEventState
// Response ID: k_EMsgDevResetEventStateResponse
// Request type: CMsgDevResetEventState
// Response type: CMsgDevResetEventStateResponse
func (d *Dota2) SendDevResetEventState(
	ctx context.Context,
	eventID protocol.EEvent,
	removeAudit bool,
) (*protocol.CMsgDevResetEventStateResponse, error) {
	req := &protocol.CMsgDevResetEventState{
		EventId:     &eventID,
		RemoveAudit: &removeAudit,
	}
	resp := &protocol.CMsgDevResetEventStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevResetEventState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevResetEventStateResponse),
		resp,
	)
}

// SendH264Unsupported sends a h 264 unsupported.
// Request ID: k_EMsgClientToGCH264Unsupported
// Request type: CMsgClientToGCH264Unsupported
func (d *Dota2) SendH264Unsupported() {
	req := &protocol.CMsgClientToGCH264Unsupported{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCH264Unsupported), req)
}

// SendHasPlayerVotedForMVP sends a has player voted for mvp.
// Request ID: k_EMsgClientToGCHasPlayerVotedForMVP
// Response ID: k_EMsgClientToGCHasPlayerVotedForMVPResponse
// Request type: CMsgClientToGCHasPlayerVotedForMVP
// Response type: CMsgClientToGCHasPlayerVotedForMVPResponse
func (d *Dota2) SendHasPlayerVotedForMVP(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCHasPlayerVotedForMVPResponse, error) {
	req := &protocol.CMsgClientToGCHasPlayerVotedForMVP{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCHasPlayerVotedForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVP),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVPResponse),
		resp,
	)
}

// SendInitialQuestionnaireResponse sends a initial questionnaire response.
// Request ID: k_EMsgGCInitialQuestionnaireResponse
// Request type: CMsgInitialQuestionnaireResponse
func (d *Dota2) SendInitialQuestionnaireResponse(
	initialSkill uint32,
) {
	req := &protocol.CMsgInitialQuestionnaireResponse{
		InitialSkill: &initialSkill,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse), req)
}

// SendInviteToGuild sends a invite to guild.
// Request ID: k_EMsgClientToGCInviteToGuild
// Response ID: k_EMsgClientToGCInviteToGuildResponse
// Request type: CMsgClientToGCInviteToGuild
// Response type: CMsgClientToGCInviteToGuildResponse
func (d *Dota2) SendInviteToGuild(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCInviteToGuild{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCInviteToGuildResponse),
		resp,
	)
}

// SendLatestConductScorecard sends a latest conduct scorecard.
// Request ID: k_EMsgClientToGCLatestConductScorecard
// Request type: CMsgPlayerConductScorecard
func (d *Dota2) SendLatestConductScorecard(
	req *protocol.CMsgPlayerConductScorecard,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard), req)
}

// SendLeagueAvailableLobbyNodes sends league available lobby nodes.
// Request ID: k_EMsgDOTALeagueAvailableLobbyNodes
// Request type: CMsgDOTALeagueAvailableLobbyNodes
func (d *Dota2) SendLeagueAvailableLobbyNodes(
	nodeInfos []*protocol.CMsgDOTALeagueAvailableLobbyNodes_NodeInfo,
) {
	req := &protocol.CMsgDOTALeagueAvailableLobbyNodes{
		NodeInfos: nodeInfos,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodes), req)
}

// SendLobbyBattleCupVictory sends a lobby battle cup victory.
// Request ID: k_EMsgLobbyBattleCupVictory
// Request type: CMsgBattleCupVictory
func (d *Dota2) SendLobbyBattleCupVictory(
	req *protocol.CMsgBattleCupVictory,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory), req)
}

// SendLobbyEventGameData sends a lobby event game data.
// Request ID: k_EMsgLobbyEventGameData
// Request type: CMsgLobbyEventGameData
func (d *Dota2) SendLobbyEventGameData(
	gameSeed uint32,
	eventWindowStartTime uint32,
) {
	req := &protocol.CMsgLobbyEventGameData{
		GameSeed:             &gameSeed,
		EventWindowStartTime: &eventWindowStartTime,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventGameData), req)
}

// SendLobbyEventGameDetails sends lobby event game details.
// Request ID: k_EMsgLobbyEventGameDetails
// Request type: CMsgLobbyEventGameDetails
func (d *Dota2) SendLobbyEventGameDetails(
	kvData []byte,
) {
	req := &protocol.CMsgLobbyEventGameDetails{
		KvData: kvData,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventGameDetails), req)
}

// SendLobbyEventPoints sends lobby event points.
// Request ID: k_EMsgLobbyEventPoints
// Request type: CMsgLobbyEventPoints
func (d *Dota2) SendLobbyEventPoints(
	eventID uint32,
	accountPoints []*protocol.CMsgLobbyEventPoints_AccountPoints,
) {
	req := &protocol.CMsgLobbyEventPoints{
		EventId:       &eventID,
		AccountPoints: accountPoints,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventPoints), req)
}

// SendLobbyGauntletProgress sends lobby gauntlet progress.
// Request ID: k_EMsgLobbyGauntletProgress
// Request type: CMsgLobbyGauntletProgress
func (d *Dota2) SendLobbyGauntletProgress(
	accounts []*protocol.CMsgLobbyGauntletProgress_AccountProgress,
) {
	req := &protocol.CMsgLobbyGauntletProgress{
		Accounts: accounts,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyGauntletProgress), req)
}

// SendLobbyPlaytestDetails sends lobby playtest details.
// Request ID: k_EMsgLobbyPlaytestDetails
// Request type: CMsgLobbyPlaytestDetails
func (d *Dota2) SendLobbyPlaytestDetails(
	jSON string,
) {
	req := &protocol.CMsgLobbyPlaytestDetails{
		Json: &jSON,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyPlaytestDetails), req)
}

// SendMMInfo sends a mm info.
// Request ID: k_EMsgClientToGCMMInfo
// Request type: CMsgClientToGCMMInfo
func (d *Dota2) SendMMInfo(
	laneSelectionFlags uint32,
	highPriorityDisabled bool,
) {
	req := &protocol.CMsgClientToGCMMInfo{
		LaneSelectionFlags:   &laneSelectionFlags,
		HighPriorityDisabled: &highPriorityDisabled,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMMInfo), req)
}

// SendManageFavorites sends manage favorites.
// Request ID: k_EMsgClientToGCManageFavorites
// Response ID: k_EMsgGCToClientManageFavoritesResponse
// Request type: CMsgClientToGCManageFavorites
// Response type: CMsgGCToClientManageFavoritesResponse
func (d *Dota2) SendManageFavorites(
	ctx context.Context,
	req *protocol.CMsgClientToGCManageFavorites,
) (*protocol.CMsgGCToClientManageFavoritesResponse, error) {
	resp := &protocol.CMsgGCToClientManageFavoritesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCManageFavorites),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientManageFavoritesResponse),
		resp,
	)
}

// SendMatchMatchmakingStats sends match matchmaking stats.
// Request ID: k_EMsgMatchMatchmakingStats
// Request type: CMsgMatchMatchmakingStats
func (d *Dota2) SendMatchMatchmakingStats(
	averageQueueTime uint32,
	maximumQueueTime uint32,
	behaviorScoreVariance protocol.EMatchBehaviorScoreVariance,
) {
	req := &protocol.CMsgMatchMatchmakingStats{
		AverageQueueTime:      &averageQueueTime,
		MaximumQueueTime:      &maximumQueueTime,
		BehaviorScoreVariance: &behaviorScoreVariance,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgMatchMatchmakingStats), req)
}

// SendMergePartyInvite sends a merge party invite.
// Request ID: k_EMsgClientToGCMergePartyInvite
// Request type: CMsgDOTAGroupMergeInvite
func (d *Dota2) SendMergePartyInvite(
	otherGroupID uint64,
) {
	req := &protocol.CMsgDOTAGroupMergeInvite{
		OtherGroupId: &otherGroupID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite), req)
}

// SendModifyGuildRole sends a modify guild role.
// Request ID: k_EMsgClientToGCModifyGuildRole
// Response ID: k_EMsgClientToGCModifyGuildRoleResponse
// Request type: CMsgClientToGCModifyGuildRole
// Response type: CMsgClientToGCModifyGuildRoleResponse
func (d *Dota2) SendModifyGuildRole(
	ctx context.Context,
	guildID uint32,
	roleID uint32,
	roleName string,
	roleFlags uint32,
) (*protocol.CMsgClientToGCModifyGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCModifyGuildRole{
		GuildId:   &guildID,
		RoleId:    &roleID,
		RoleName:  &roleName,
		RoleFlags: &roleFlags,
	}
	resp := &protocol.CMsgClientToGCModifyGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCModifyGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCModifyGuildRoleResponse),
		resp,
	)
}

// SendNeutralItemStats sends neutral item stats.
// Request ID: k_EMsgNeutralItemStats
// Request type: CMsgNeutralItemStats
func (d *Dota2) SendNeutralItemStats(
	neutralItems []*protocol.CMsgNeutralItemStats_NeutralItem,
) {
	req := &protocol.CMsgNeutralItemStats{
		NeutralItems: neutralItems,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgNeutralItemStats), req)
}

// SendOverwatchReplayError sends a overwatch replay error.
// Request ID: k_EMsgClientToGCOverwatchReplayError
// Request type: CMsgClientToGCOverwatchReplayError
func (d *Dota2) SendOverwatchReplayError(
	overwatchReplayID uint64,
) {
	req := &protocol.CMsgClientToGCOverwatchReplayError{
		OverwatchReplayId: &overwatchReplayID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverwatchReplayError), req)
}

// SendPartyReadyCheck sends a party ready check.
// Request ID: k_EMsgPartyReadyCheckRequest
// Response ID: k_EMsgPartyReadyCheckResponse
// Request type: CMsgPartyReadyCheckRequest
// Response type: CMsgPartyReadyCheckResponse
func (d *Dota2) SendPartyReadyCheck(
	ctx context.Context,
) (*protocol.CMsgPartyReadyCheckResponse, error) {
	req := &protocol.CMsgPartyReadyCheckRequest{}
	resp := &protocol.CMsgPartyReadyCheckResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckResponse),
		resp,
	)
}

// SendPeriodicResourceUpdated sends a periodic resource updated.
// Request ID: k_EMsgDOTAPeriodicResourceUpdated
// Request type: CMsgDOTAPeriodicResourceUpdated
func (d *Dota2) SendPeriodicResourceUpdated(
	req *protocol.CMsgDOTAPeriodicResourceUpdated,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTAPeriodicResourceUpdated), req)
}

// SendPingData sends a ping data.
// Request ID: k_EMsgClientToGCPingData
// Request type: CMsgClientPingData
func (d *Dota2) SendPingData(
	relayCodes []uint32,
	relayPings []uint32,
	regionCodes []uint32,
	regionPings []uint32,
	regionPingFailedBitmask uint32,
) {
	req := &protocol.CMsgClientPingData{
		RelayCodes:              relayCodes,
		RelayPings:              relayPings,
		RegionCodes:             regionCodes,
		RegionPings:             regionPings,
		RegionPingFailedBitmask: &regionPingFailedBitmask,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPingData), req)
}

// SendProfileUpdate sends a profile update.
// Request ID: k_EMsgProfileUpdate
// Response ID: k_EMsgProfileUpdateResponse
// Request type: CMsgProfileUpdate
// Response type: CMsgProfileUpdateResponse
func (d *Dota2) SendProfileUpdate(
	ctx context.Context,
	backgroundItemID uint64,
	featuredHeroIDs []uint32,
) (*protocol.CMsgProfileUpdateResponse, error) {
	req := &protocol.CMsgProfileUpdate{
		BackgroundItemId: &backgroundItemID,
		FeaturedHeroIds:  featuredHeroIDs,
	}
	resp := &protocol.CMsgProfileUpdateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileUpdate),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileUpdateResponse),
		resp,
	)
}

// SendReadyUp sends a ready up.
// Request ID: k_EMsgGCReadyUp
// Request type: CMsgReadyUp
func (d *Dota2) SendReadyUp(
	req *protocol.CMsgReadyUp,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCReadyUp), req)
}

// SendRecalibrateMMR sends a recalibrate mmr.
// Request ID: k_EMsgClientToGCRecalibrateMMR
// Response ID: k_EMsgClientToGCRecalibrateMMRResponse
// Request type: CMsgClientToGCRecalibrateMMR
// Response type: CMsgClientToGCRecalibrateMMRResponse
func (d *Dota2) SendRecalibrateMMR(
	ctx context.Context,
) (*protocol.CMsgClientToGCRecalibrateMMRResponse, error) {
	req := &protocol.CMsgClientToGCRecalibrateMMR{}
	resp := &protocol.CMsgClientToGCRecalibrateMMRResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecalibrateMMR),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecalibrateMMRResponse),
		resp,
	)
}

// SendRemoveFilteredPlayer sends a remove filtered player.
// Request ID: k_EMsgClientToGCRemoveFilteredPlayer
// Response ID: k_EMsgGCToClientRemoveFilteredPlayerResponse
// Request type: CMsgClientToGCRemoveFilteredPlayer
// Response type: CMsgGCToClientRemoveFilteredPlayerResponse
func (d *Dota2) SendRemoveFilteredPlayer(
	ctx context.Context,
	accountIDToRemove uint32,
) (*protocol.CMsgGCToClientRemoveFilteredPlayerResponse, error) {
	req := &protocol.CMsgClientToGCRemoveFilteredPlayer{
		AccountIdToRemove: &accountIDToRemove,
	}
	resp := &protocol.CMsgGCToClientRemoveFilteredPlayerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveFilteredPlayer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRemoveFilteredPlayerResponse),
		resp,
	)
}

// SendRemoveGuildRole sends a remove guild role.
// Request ID: k_EMsgClientToGCRemoveGuildRole
// Response ID: k_EMsgClientToGCRemoveGuildRoleResponse
// Request type: CMsgClientToGCRemoveGuildRole
// Response type: CMsgClientToGCRemoveGuildRoleResponse
func (d *Dota2) SendRemoveGuildRole(
	ctx context.Context,
	guildID uint32,
	roleID uint32,
) (*protocol.CMsgClientToGCRemoveGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCRemoveGuildRole{
		GuildId: &guildID,
		RoleId:  &roleID,
	}
	resp := &protocol.CMsgClientToGCRemoveGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveGuildRoleResponse),
		resp,
	)
}

// SendSpectatorLobbyGameDetails sends spectator lobby game details.
// Request ID: k_EMsgSpectatorLobbyGameDetails
// Request type: CMsgSpectatorLobbyGameDetails
func (d *Dota2) SendSpectatorLobbyGameDetails(
	req *protocol.CMsgSpectatorLobbyGameDetails,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgSpectatorLobbyGameDetails), req)
}

// SendTeamInvite_GCResponseToInvitee sends a team invite _ gc response to invitee.
// Request ID: k_EMsgGCTeamInvite_GCResponseToInvitee
// Request type: CMsgDOTATeamInvite_GCResponseToInvitee
func (d *Dota2) SendTeamInvite_GCResponseToInvitee(
	result protocol.ETeamInviteResult,
	teamName string,
) {
	req := &protocol.CMsgDOTATeamInvite_GCResponseToInvitee{
		Result:   &result,
		TeamName: &teamName,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee), req)
}

// SendUnderDraftBuy sends a under draft buy.
// Request ID: k_EMsgClientToGCUnderDraftBuy
// Response ID: k_EMsgClientToGCUnderDraftBuyResponse
// Request type: CMsgClientToGCUnderDraftBuy
// Response type: CMsgClientToGCUnderDraftBuyResponse
func (d *Dota2) SendUnderDraftBuy(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
) (*protocol.CMsgClientToGCUnderDraftBuyResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftBuy{
		EventId: &eventID,
		SlotId:  &slotID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftBuyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftBuy),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftBuyResponse),
		resp,
	)
}

// SendUnderDraftRollBackBench sends a under draft roll back bench.
// Request ID: k_EMsgClientToGCUnderDraftRollBackBench
// Response ID: k_EMsgClientToGCUnderDraftRollBackBenchResponse
// Request type: CMsgClientToGCUnderDraftRollBackBench
// Response type: CMsgClientToGCUnderDraftRollBackBenchResponse
func (d *Dota2) SendUnderDraftRollBackBench(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftRollBackBenchResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRollBackBench{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRollBackBenchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRollBackBench),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRollBackBenchResponse),
		resp,
	)
}

// SendUnderDraftSell sends a under draft sell.
// Request ID: k_EMsgClientToGCUnderDraftSell
// Response ID: k_EMsgClientToGCUnderDraftSellResponse
// Request type: CMsgClientToGCUnderDraftSell
// Response type: CMsgClientToGCUnderDraftSellResponse
func (d *Dota2) SendUnderDraftSell(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
) (*protocol.CMsgClientToGCUnderDraftSellResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftSell{
		EventId: &eventID,
		SlotId:  &slotID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftSellResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftSell),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftSellResponse),
		resp,
	)
}

// SendUpdatePartyBeacon sends a update party beacon.
// Request ID: k_EMsgClientToGCUpdatePartyBeacon
// Request type: CMsgClientToGCUpdatePartyBeacon
func (d *Dota2) SendUpdatePartyBeacon(
	action protocol.CMsgClientToGCUpdatePartyBeacon_Action,
) {
	req := &protocol.CMsgClientToGCUpdatePartyBeacon{
		Action: &action,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUpdatePartyBeacon), req)
}

// SendVerifyFavoritePlayers sends verify favorite players.
// Request ID: k_EMsgClientToGCVerifyFavoritePlayers
// Response ID: k_EMsgGCToClientVerifyFavoritePlayersResponse
// Request type: CMsgClientToGCVerifyFavoritePlayers
// Response type: CMsgGCToClientVerifyFavoritePlayersResponse
func (d *Dota2) SendVerifyFavoritePlayers(
	ctx context.Context,
	accountIDs []uint32,
) (*protocol.CMsgGCToClientVerifyFavoritePlayersResponse, error) {
	req := &protocol.CMsgClientToGCVerifyFavoritePlayers{
		AccountIds: accountIDs,
	}
	resp := &protocol.CMsgGCToClientVerifyFavoritePlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVerifyFavoritePlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientVerifyFavoritePlayersResponse),
		resp,
	)
}

// SendWatchGame sends a watch game.
// Request ID: k_EMsgGCWatchGame
// Response ID: k_EMsgGCWatchGameResponse
// Request type: CMsgWatchGame
// Response type: CMsgWatchGameResponse
func (d *Dota2) SendWatchGame(
	ctx context.Context,
	serverSteamid steamid.SteamId,
	watchServerSteamid steamid.SteamId,
	lobbyID uint64,
	regions []uint32,
) (*protocol.CMsgWatchGameResponse, error) {
	serverSteamidU64Val := uint64(serverSteamid)
	serverSteamidU64 := &serverSteamidU64Val
	watchServerSteamidU64Val := uint64(watchServerSteamid)
	watchServerSteamidU64 := &watchServerSteamidU64Val
	req := &protocol.CMsgWatchGame{
		ServerSteamid:      serverSteamidU64,
		WatchServerSteamid: watchServerSteamidU64,
		LobbyId:            &lobbyID,
		Regions:            regions,
	}
	resp := &protocol.CMsgWatchGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchGame),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchGameResponse),
		resp,
	)
}

// SetCompendiumSelection sets a compendium selection.
// Request ID: k_EMsgGCCompendiumSetSelection
// Response ID: k_EMsgGCCompendiumSetSelectionResponse
// Request type: CMsgDOTACompendiumSelection
// Response type: CMsgDOTACompendiumSelectionResponse
func (d *Dota2) SetCompendiumSelection(
	ctx context.Context,
	selectionIndex uint32,
	selection uint32,
	leagueid uint32,
) (*protocol.CMsgDOTACompendiumSelectionResponse, error) {
	req := &protocol.CMsgDOTACompendiumSelection{
		SelectionIndex: &selectionIndex,
		Selection:      &selection,
		Leagueid:       &leagueid,
	}
	resp := &protocol.CMsgDOTACompendiumSelectionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse),
		resp,
	)
}

// SetDPCFavoriteState sets a dpc favorite state.
// Request ID: k_EMsgClientToGCSetDPCFavoriteState
// Response ID: k_EMsgClientToGCSetDPCFavoriteStateResponse
// Request type: CMsgClientToGCSetDPCFavoriteState
// Response type: CMsgClientToGCSetDPCFavoriteStateResponse
func (d *Dota2) SetDPCFavoriteState(
	ctx context.Context,
	favoriteType protocol.EDPCFavoriteType,
	favoriteID uint32,
	enabled bool,
) (*protocol.CMsgClientToGCSetDPCFavoriteStateResponse, error) {
	req := &protocol.CMsgClientToGCSetDPCFavoriteState{
		FavoriteType: &favoriteType,
		FavoriteId:   &favoriteID,
		Enabled:      &enabled,
	}
	resp := &protocol.CMsgClientToGCSetDPCFavoriteStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetDPCFavoriteState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetDPCFavoriteStateResponse),
		resp,
	)
}

// SetEventActiveSeasonID sets a event active season id.
// Request ID: k_EMsgClientToGCSetEventActiveSeasonID
// Response ID: k_EMsgClientToGCSetEventActiveSeasonIDResponse
// Request type: CMsgClientToGCSetEventActiveSeasonID
// Response type: CMsgClientToGCSetEventActiveSeasonIDResponse
func (d *Dota2) SetEventActiveSeasonID(
	ctx context.Context,
	eventID uint32,
	activeSeasonID uint32,
) (*protocol.CMsgClientToGCSetEventActiveSeasonIDResponse, error) {
	req := &protocol.CMsgClientToGCSetEventActiveSeasonID{
		EventId:        &eventID,
		ActiveSeasonId: &activeSeasonID,
	}
	resp := &protocol.CMsgClientToGCSetEventActiveSeasonIDResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetEventActiveSeasonID),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetEventActiveSeasonIDResponse),
		resp,
	)
}

// SetFavoriteTeam sets a favorite team.
// Request ID: k_EMsgDOTASetFavoriteTeam
// Request type: CMsgDOTASetFavoriteTeam
func (d *Dota2) SetFavoriteTeam(
	teamID uint32,
	eventID uint32,
) {
	req := &protocol.CMsgDOTASetFavoriteTeam{
		TeamId:  &teamID,
		EventId: &eventID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTASetFavoriteTeam), req)
}

// SetGuildInfo sets a guild info.
// Request ID: k_EMsgClientToGCSetGuildInfo
// Response ID: k_EMsgClientToGCSetGuildInfoResponse
// Request type: CMsgClientToGCSetGuildInfo
// Response type: CMsgClientToGCSetGuildInfoResponse
func (d *Dota2) SetGuildInfo(
	ctx context.Context,
	req *protocol.CMsgClientToGCSetGuildInfo,
) (*protocol.CMsgClientToGCSetGuildInfoResponse, error) {
	resp := &protocol.CMsgClientToGCSetGuildInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildInfoResponse),
		resp,
	)
}

// SetGuildMemberRole sets a guild member role.
// Request ID: k_EMsgClientToGCSetGuildMemberRole
// Response ID: k_EMsgClientToGCSetGuildMemberRoleResponse
// Request type: CMsgClientToGCSetGuildMemberRole
// Response type: CMsgClientToGCSetGuildMemberRoleResponse
func (d *Dota2) SetGuildMemberRole(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
	targetRoleID uint32,
) (*protocol.CMsgClientToGCSetGuildMemberRoleResponse, error) {
	req := &protocol.CMsgClientToGCSetGuildMemberRole{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
		TargetRoleId:    &targetRoleID,
	}
	resp := &protocol.CMsgClientToGCSetGuildMemberRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildMemberRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildMemberRoleResponse),
		resp,
	)
}

// SetGuildRoleOrder sets a guild role order.
// Request ID: k_EMsgClientToGCSetGuildRoleOrder
// Response ID: k_EMsgClientToGCSetGuildRoleOrderResponse
// Request type: CMsgClientToGCSetGuildRoleOrder
// Response type: CMsgClientToGCSetGuildRoleOrderResponse
func (d *Dota2) SetGuildRoleOrder(
	ctx context.Context,
	guildID uint32,
	requestedRoleIDs []uint32,
	previousRoleIDs []uint32,
) (*protocol.CMsgClientToGCSetGuildRoleOrderResponse, error) {
	req := &protocol.CMsgClientToGCSetGuildRoleOrder{
		GuildId:          &guildID,
		RequestedRoleIds: requestedRoleIDs,
		PreviousRoleIds:  previousRoleIDs,
	}
	resp := &protocol.CMsgClientToGCSetGuildRoleOrderResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildRoleOrder),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildRoleOrderResponse),
		resp,
	)
}

// SetLobbyCoach sets a lobby coach.
// Request ID: k_EMsgGCPracticeLobbySetCoach
// Request type: CMsgPracticeLobbySetCoach
func (d *Dota2) SetLobbyCoach(
	team protocol.DOTA_GC_TEAM,
) {
	req := &protocol.CMsgPracticeLobbySetCoach{
		Team: &team,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbySetCoach), req)
}

// SetLobbyDetails sets lobby details.
// Request ID: k_EMsgGCPracticeLobbySetDetails
// Request type: CMsgPracticeLobbySetDetails
func (d *Dota2) SetLobbyDetails(
	req *protocol.CMsgPracticeLobbySetDetails,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails), req)
}

// SetMatchHistoryAccess sets match history access.
// Request ID: k_EMsgGCSetMatchHistoryAccess
// Response ID: k_EMsgGCSetMatchHistoryAccessResponse
// Request type: CMsgDOTASetMatchHistoryAccess
// Response type: CMsgDOTASetMatchHistoryAccessResponse
func (d *Dota2) SetMatchHistoryAccess(
	ctx context.Context,
	allow3RdPartyMatchHistory bool,
) (*protocol.CMsgDOTASetMatchHistoryAccessResponse, error) {
	req := &protocol.CMsgDOTASetMatchHistoryAccess{
		Allow_3RdPartyMatchHistory: &allow3RdPartyMatchHistory,
	}
	resp := &protocol.CMsgDOTASetMatchHistoryAccessResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccess),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccessResponse),
		resp,
	)
}

// SetMemberPartyCoach sets a member party coach.
// Request ID: k_EMsgGCPartyMemberSetCoach
// Request type: CMsgDOTAPartyMemberSetCoach
func (d *Dota2) SetMemberPartyCoach(
	wantsCoach bool,
) {
	req := &protocol.CMsgDOTAPartyMemberSetCoach{
		WantsCoach: &wantsCoach,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPartyMemberSetCoach), req)
}

// SetPartyBuilderOptions sets party builder options.
// Request ID: k_EMsgClientToGCSetPartyBuilderOptions
// Request type: CMsgPartyBuilderOptions
func (d *Dota2) SetPartyBuilderOptions(
	additionalSlots uint32,
	matchType protocol.MatchType,
	matchgroups uint32,
	language protocol.MatchLanguages,
) {
	req := &protocol.CMsgPartyBuilderOptions{
		AdditionalSlots: &additionalSlots,
		MatchType:       &matchType,
		Matchgroups:     &matchgroups,
		Language:        &language,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions), req)
}

// SetPartyLeader sets a party leader.
// Request ID: k_EMsgClientToGCSetPartyLeader
// Request type: CMsgDOTASetGroupLeader
func (d *Dota2) SetPartyLeader(
	newLeaderSteamid steamid.SteamId,
) {
	newLeaderSteamidU64Val := uint64(newLeaderSteamid)
	newLeaderSteamidU64 := &newLeaderSteamidU64Val
	req := &protocol.CMsgDOTASetGroupLeader{
		NewLeaderSteamid: newLeaderSteamidU64,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader), req)
}

// SetPartyOpen sets a party open.
// Request ID: k_EMsgClientToGCSetPartyOpen
// Request type: CMsgDOTASetGroupOpenStatus
func (d *Dota2) SetPartyOpen(
	open bool,
) {
	req := &protocol.CMsgDOTASetGroupOpenStatus{
		Open: &open,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen), req)
}

// SetProfileCardSlots sets profile card slots.
// Request ID: k_EMsgClientToGCSetProfileCardSlots
// Request type: CMsgClientToGCSetProfileCardSlots
func (d *Dota2) SetProfileCardSlots(
	slots []*protocol.CMsgClientToGCSetProfileCardSlots_CardSlot,
) {
	req := &protocol.CMsgClientToGCSetProfileCardSlots{
		Slots: slots,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetProfileCardSlots), req)
}

// SetProfilePrivacy sets a profile privacy.
// Request ID: k_EMsgGCSetProfilePrivacy
// Response ID: k_EMsgGCSetProfilePrivacyResponse
// Request type: CMsgDOTASetProfilePrivacy
// Response type: CMsgDOTASetProfilePrivacyResponse
func (d *Dota2) SetProfilePrivacy(
	ctx context.Context,
	profilePrivate bool,
) (*protocol.CMsgDOTASetProfilePrivacyResponse, error) {
	req := &protocol.CMsgDOTASetProfilePrivacy{
		ProfilePrivate: &profilePrivate,
	}
	resp := &protocol.CMsgDOTASetProfilePrivacyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetProfilePrivacy),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetProfilePrivacyResponse),
		resp,
	)
}

// SetSpectatorLobbyDetails sets spectator lobby details.
// Request ID: k_EMsgClientToGCSetSpectatorLobbyDetails
// Request type: CMsgSetSpectatorLobbyDetails
func (d *Dota2) SetSpectatorLobbyDetails(
	req *protocol.CMsgSetSpectatorLobbyDetails,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetails), req)
}

// SpectateFriendGame spectates a friend game.
// Request ID: k_EMsgGCSpectateFriendGame
// Response ID: k_EMsgGCSpectateFriendGameResponse
// Request type: CMsgSpectateFriendGame
// Response type: CMsgSpectateFriendGameResponse
func (d *Dota2) SpectateFriendGame(
	ctx context.Context,
	steamID steamid.SteamId,
	live bool,
) (*protocol.CMsgSpectateFriendGameResponse, error) {
	steamIDU64Val := uint64(steamID)
	steamIDU64 := &steamIDU64Val
	req := &protocol.CMsgSpectateFriendGame{
		SteamId: steamIDU64,
		Live:    &live,
	}
	resp := &protocol.CMsgSpectateFriendGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSpectateFriendGame),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSpectateFriendGameResponse),
		resp,
	)
}

// StartFindingMatch starts a finding match.
// Request ID: k_EMsgGCStartFindingMatch
// Response ID: k_EMsgGCStartFindingMatchResponse
// Request type: CMsgStartFindingMatch
// Response type: CMsgStartFindingMatchResult
func (d *Dota2) StartFindingMatch(
	ctx context.Context,
	req *protocol.CMsgStartFindingMatch,
) (*protocol.CMsgStartFindingMatchResult, error) {
	resp := &protocol.CMsgStartFindingMatchResult{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCStartFindingMatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCStartFindingMatchResponse),
		resp,
	)
}

// StartTriviaSession starts a trivia session.
// Request ID: k_EMsgStartTriviaSession
// Response ID: k_EMsgStartTriviaSessionResponse
// Request type: CMsgDOTAStartTriviaSession
// Response type: CMsgDOTAStartTriviaSessionResponse
func (d *Dota2) StartTriviaSession(
	ctx context.Context,
) (*protocol.CMsgDOTAStartTriviaSessionResponse, error) {
	req := &protocol.CMsgDOTAStartTriviaSession{}
	resp := &protocol.CMsgDOTAStartTriviaSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgStartTriviaSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgStartTriviaSessionResponse),
		resp,
	)
}

// StartWatchingOverwatch starts a watching overwatch.
// Request ID: k_EMsgClientToGCStartWatchingOverwatch
// Request type: CMsgClientToGCStartWatchingOverwatch
func (d *Dota2) StartWatchingOverwatch(
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
) {
	req := &protocol.CMsgClientToGCStartWatchingOverwatch{
		OverwatchReplayId: &overwatchReplayID,
		TargetPlayerSlot:  &targetPlayerSlot,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCStartWatchingOverwatch), req)
}

// StopFindingMatch stops a finding match.
// Request ID: k_EMsgGCStopFindingMatch
// Request type: CMsgStopFindingMatch
func (d *Dota2) StopFindingMatch(
	acceptCooldown bool,
) {
	req := &protocol.CMsgStopFindingMatch{
		AcceptCooldown: &acceptCooldown,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCStopFindingMatch), req)
}

// StopWatchingOverwatch stops a watching overwatch.
// Request ID: k_EMsgClientToGCStopWatchingOverwatch
// Request type: CMsgClientToGCStopWatchingOverwatch
func (d *Dota2) StopWatchingOverwatch(
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
) {
	req := &protocol.CMsgClientToGCStopWatchingOverwatch{
		OverwatchReplayId: &overwatchReplayID,
		TargetPlayerSlot:  &targetPlayerSlot,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCStopWatchingOverwatch), req)
}

// SubmitCoachTeammateRating submits a coach teammate rating.
// Request ID: k_EMsgClientToGCSubmitCoachTeammateRating
// Response ID: k_EMsgClientToGCSubmitCoachTeammateRatingResponse
// Request type: CMsgClientToGCSubmitCoachTeammateRating
// Response type: CMsgClientToGCSubmitCoachTeammateRatingResponse
func (d *Dota2) SubmitCoachTeammateRating(
	ctx context.Context,
	matchID uint64,
	coachAccountID uint32,
	rating protocol.ECoachTeammateRating,
	reason string,
) (*protocol.CMsgClientToGCSubmitCoachTeammateRatingResponse, error) {
	req := &protocol.CMsgClientToGCSubmitCoachTeammateRating{
		MatchId:        &matchID,
		CoachAccountId: &coachAccountID,
		Rating:         &rating,
		Reason:         &reason,
	}
	resp := &protocol.CMsgClientToGCSubmitCoachTeammateRatingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitCoachTeammateRating),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitCoachTeammateRatingResponse),
		resp,
	)
}

// SubmitDraftTriviaMatchAnswer submits a draft trivia match answer.
// Request ID: k_EMsgClientToGCSubmitDraftTriviaMatchAnswer
// Response ID: k_EMsgClientToGCSubmitDraftTriviaMatchAnswerResponse
// Request type: CMsgClientToGCSubmitDraftTriviaMatchAnswer
// Response type: CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse
func (d *Dota2) SubmitDraftTriviaMatchAnswer(
	ctx context.Context,
	choseRadiantAsWinner bool,
	eventID uint32,
	endTime uint32,
) (*protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse, error) {
	req := &protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswer{
		ChoseRadiantAsWinner: &choseRadiantAsWinner,
		EventId:              &eventID,
		EndTime:              &endTime,
	}
	resp := &protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswerResponse),
		resp,
	)
}

// SubmitInfoPlayer submits a info player.
// Request ID: k_EMsgGCPlayerInfoSubmit
// Response ID: k_EMsgGCPlayerInfoSubmitResponse
// Request type: CMsgGCPlayerInfoSubmit
// Response type: CMsgGCPlayerInfoSubmitResponse
func (d *Dota2) SubmitInfoPlayer(
	ctx context.Context,
	name string,
	countryCode string,
	fantasyRole uint32,
	teamID uint32,
	sponsor string,
) (*protocol.CMsgGCPlayerInfoSubmitResponse, error) {
	req := &protocol.CMsgGCPlayerInfoSubmit{
		Name:        &name,
		CountryCode: &countryCode,
		FantasyRole: &fantasyRole,
		TeamId:      &teamID,
		Sponsor:     &sponsor,
	}
	resp := &protocol.CMsgGCPlayerInfoSubmitResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmit),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmitResponse),
		resp,
	)
}

// SubmitLobbyMVPVote submits a lobby mvp vote.
// Request ID: k_EMsgGCSubmitLobbyMVPVote
// Response ID: k_EMsgGCSubmitLobbyMVPVoteResponse
// Request type: CMsgDOTASubmitLobbyMVPVote
// Response type: CMsgDOTASubmitLobbyMVPVoteResponse
func (d *Dota2) SubmitLobbyMVPVote(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgDOTASubmitLobbyMVPVoteResponse, error) {
	req := &protocol.CMsgDOTASubmitLobbyMVPVote{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgDOTASubmitLobbyMVPVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVote),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVoteResponse),
		resp,
	)
}

// SubmitOWConviction submits a ow conviction.
// Request ID: k_EMsgClientToGCSubmitOWConviction
// Response ID: k_EMsgClientToGCSubmitOWConvictionResponse
// Request type: CMsgClientToGCSubmitOWConviction
// Response type: CMsgClientToGCSubmitOWConvictionResponse
func (d *Dota2) SubmitOWConviction(
	ctx context.Context,
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
	cheatingConviction protocol.EOverwatchConviction,
	griefingConviction protocol.EOverwatchConviction,
) (*protocol.CMsgClientToGCSubmitOWConvictionResponse, error) {
	req := &protocol.CMsgClientToGCSubmitOWConviction{
		OverwatchReplayId:  &overwatchReplayID,
		TargetPlayerSlot:   &targetPlayerSlot,
		CheatingConviction: &cheatingConviction,
		GriefingConviction: &griefingConviction,
	}
	resp := &protocol.CMsgClientToGCSubmitOWConvictionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitOWConviction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitOWConvictionResponse),
		resp,
	)
}

// SubmitPlayerMatchSurvey submits a player match survey.
// Request ID: k_EMsgClientToGCSubmitPlayerMatchSurvey
// Response ID: k_EMsgClientToGCSubmitPlayerMatchSurveyResponse
// Request type: CMsgClientToGCSubmitPlayerMatchSurvey
// Response type: CMsgClientToGCSubmitPlayerMatchSurveyResponse
func (d *Dota2) SubmitPlayerMatchSurvey(
	ctx context.Context,
	matchID uint64,
	rating int32,
	flags uint32,
) (*protocol.CMsgClientToGCSubmitPlayerMatchSurveyResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPlayerMatchSurvey{
		MatchId: &matchID,
		Rating:  &rating,
		Flags:   &flags,
	}
	resp := &protocol.CMsgClientToGCSubmitPlayerMatchSurveyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPlayerMatchSurvey),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPlayerMatchSurveyResponse),
		resp,
	)
}

// SubmitPlayerReport submits a player report.
// Request ID: k_EMsgGCSubmitPlayerReport
// Response ID: k_EMsgGCSubmitPlayerReportResponse
// Request type: CMsgDOTASubmitPlayerReport
// Response type: CMsgDOTASubmitPlayerReportResponse
func (d *Dota2) SubmitPlayerReport(
	ctx context.Context,
	targetAccountID uint32,
	reportFlags uint32,
	lobbyID uint64,
	comment string,
) (*protocol.CMsgDOTASubmitPlayerReportResponse, error) {
	req := &protocol.CMsgDOTASubmitPlayerReport{
		TargetAccountId: &targetAccountID,
		ReportFlags:     &reportFlags,
		LobbyId:         &lobbyID,
		Comment:         &comment,
	}
	resp := &protocol.CMsgDOTASubmitPlayerReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReport),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReportResponse),
		resp,
	)
}

// SubmitPrivateCoachingSessionRating submits a private coaching session rating.
// Request ID: k_EMsgClientToGCSubmitPrivateCoachingSessionRating
// Response ID: k_EMsgClientToGCSubmitPrivateCoachingSessionRatingResponse
// Request type: CMsgClientToGCSubmitPrivateCoachingSessionRating
// Response type: CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse
func (d *Dota2) SubmitPrivateCoachingSessionRating(
	ctx context.Context,
	coachingSessionID uint64,
	sessionRating protocol.ECoachTeammateRating,
) (*protocol.CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPrivateCoachingSessionRating{
		CoachingSessionId: &coachingSessionID,
		SessionRating:     &sessionRating,
	}
	resp := &protocol.CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPrivateCoachingSessionRating),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPrivateCoachingSessionRatingResponse),
		resp,
	)
}

// SubmitTriviaQuestionAnswer submits a trivia question answer.
// Request ID: k_EMsgSubmitTriviaQuestionAnswer
// Response ID: k_EMsgSubmitTriviaQuestionAnswerResponse
// Request type: CMsgDOTASubmitTriviaQuestionAnswer
// Response type: CMsgDOTASubmitTriviaQuestionAnswerResponse
func (d *Dota2) SubmitTriviaQuestionAnswer(
	ctx context.Context,
	questionID uint32,
	answerIndex uint32,
) (*protocol.CMsgDOTASubmitTriviaQuestionAnswerResponse, error) {
	req := &protocol.CMsgDOTASubmitTriviaQuestionAnswer{
		QuestionId:  &questionID,
		AnswerIndex: &answerIndex,
	}
	resp := &protocol.CMsgDOTASubmitTriviaQuestionAnswerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswerResponse),
		resp,
	)
}

// ToggleLobbyBroadcastChannelCameramanStatus toggles lobby broadcast channel cameraman status.
// Request ID: k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus
// Request type: CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus
func (d *Dota2) ToggleLobbyBroadcastChannelCameramanStatus() {
	req := &protocol.CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus), req)
}

// TransferTeamAdmin transfers a team admin.
// Request ID: k_EMsgGCTransferTeamAdmin
// Response ID: k_EMsgGCTransferTeamAdminResponse
// Request type: CMsgDOTATransferTeamAdmin
// Response type: CMsgDOTATransferTeamAdminResponse
func (d *Dota2) TransferTeamAdmin(
	ctx context.Context,
	newAdminAccountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTATransferTeamAdminResponse, error) {
	req := &protocol.CMsgDOTATransferTeamAdmin{
		NewAdminAccountId: &newAdminAccountID,
		TeamId:            &teamID,
	}
	resp := &protocol.CMsgDOTATransferTeamAdminResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTransferTeamAdmin),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTransferTeamAdminResponse),
		resp,
	)
}

// UpgradeLeagueItem upgrades a league item.
// Request ID: k_EMsgUpgradeLeagueItem
// Response ID: k_EMsgUpgradeLeagueItemResponse
// Request type: CMsgUpgradeLeagueItem
// Response type: CMsgUpgradeLeagueItemResponse
func (d *Dota2) UpgradeLeagueItem(
	ctx context.Context,
	matchID uint64,
	leagueID uint32,
) (*protocol.CMsgUpgradeLeagueItemResponse, error) {
	req := &protocol.CMsgUpgradeLeagueItem{
		MatchId:  &matchID,
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgUpgradeLeagueItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgUpgradeLeagueItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgUpgradeLeagueItemResponse),
		resp,
	)
}

// VoteForArcana votes a for arcana.
// Request ID: k_EMsgClientToGCVoteForArcana
// Response ID: k_EMsgClientToGCVoteForArcanaResponse
// Request type: CMsgClientToGCVoteForArcana
// Response type: CMsgClientToGCVoteForArcanaResponse
func (d *Dota2) VoteForArcana(
	ctx context.Context,
	matches []*protocol.CMsgArcanaVoteMatchVotes,
) (*protocol.CMsgClientToGCVoteForArcanaResponse, error) {
	req := &protocol.CMsgClientToGCVoteForArcana{
		Matches: matches,
	}
	resp := &protocol.CMsgClientToGCVoteForArcanaResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForArcana),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForArcanaResponse),
		resp,
	)
}

// VoteForMVP votes a for mvp.
// Request ID: k_EMsgClientToGCVoteForMVP
// Response ID: k_EMsgClientToGCVoteForMVPResponse
// Request type: CMsgClientToGCVoteForMVP
// Response type: CMsgClientToGCVoteForMVPResponse
func (d *Dota2) VoteForMVP(
	ctx context.Context,
	matchID uint64,
	accountID uint32,
) (*protocol.CMsgClientToGCVoteForMVPResponse, error) {
	req := &protocol.CMsgClientToGCVoteForMVP{
		MatchId:   &matchID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCVoteForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForMVP),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForMVPResponse),
		resp,
	)
}

// VoteMVPTimeout votes a mvp timeout.
// Request ID: k_EMsgClientToGCMVPVoteTimeout
// Response ID: k_EMsgClientToGCMVPVoteTimeoutResponse
// Request type: CMsgClientToGCMVPVoteTimeout
// Response type: CMsgClientToGCMVPVoteTimeoutResponse
func (d *Dota2) VoteMVPTimeout(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCMVPVoteTimeoutResponse, error) {
	req := &protocol.CMsgClientToGCMVPVoteTimeout{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCMVPVoteTimeoutResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMVPVoteTimeout),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMVPVoteTimeoutResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Dota2) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientAccountGuildEventDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.AccountGuildEventDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildChallengeUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ActiveGuildChallengeUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildContractsUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ActiveGuildContractsUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientArcanaVotesUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.ArcanaVotesUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupListRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupListRequest{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupRequest{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCBroadcastNotification)] = d.getEventEmitter(func() events.Event {
		return &events.BroadcastNotification{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapPathCompleted)] = d.getEventEmitter(func() events.Event {
		return &events.CavernCrawlMapPathCompleted{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.CavernCrawlMapUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCChatModeratorBan)] = d.getEventEmitter(func() events.Event {
		return &events.ChatModeratorBan{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled)] = d.getEventEmitter(func() events.Event {
		return &events.ChatRegionsEnabled{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientClaimEventActionUsingItemCompleted)] = d.getEventEmitter(func() events.Event {
		return &events.ClaimEventActionUsingItemCompleted{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCClientSuspended)] = d.getEventEmitter(func() events.Event {
		return &events.ClientSuspended{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCoachTeammateRatingsChanged)] = d.getEventEmitter(func() events.Event {
		return &events.CoachTeammateRatingsChanged{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCommendNotification)] = d.getEventEmitter(func() events.Event {
		return &events.CommendNotification{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule)] = d.getEventEmitter(func() events.Event {
		return &events.DOTAWeekendTourneySchedule{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientEmoticonData)] = d.getEventEmitter(func() events.Event {
		return &events.EmoticonData{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyFinalPlayerStats{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildFeedUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildFeedUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembersDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildMembersDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembershipUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildMembershipUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientHeroStatueCreateResult)] = d.getEventEmitter(func() events.Event {
		return &events.HeroStatueCreateResult{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue)] = d.getEventEmitter(func() events.Event {
		return &events.KickedFromMatchmakingQueue{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCLeagueAdminList)] = d.getEventEmitter(func() events.Event {
		return &events.LeagueAdminList{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyMVPAwarded{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCLobbyUpdateBroadcastChannelInfo)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyUpdateBroadcastChannelInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMatchGroupsVersion)] = d.getEventEmitter(func() events.Event {
		return &events.MatchGroupsVersion{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMatchSignedOut)] = d.getEventEmitter(func() events.Event {
		return &events.MatchSignedOut{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergeGroupInviteReply{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergePartyResponseReply{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientNotificationsUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.NotificationsUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCNotifyAccountFlagsChange)] = d.getEventEmitter(func() events.Event {
		return &events.NotifyAccountFlagsChange{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientOverwatchCasesAvailable)] = d.getEventEmitter(func() events.Event {
		return &events.OverwatchCasesAvailable{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartyBeaconUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.PartyBeaconUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCPartyLeaderWatchGamePrompt)] = d.getEventEmitter(func() events.Event {
		return &events.PartyLeaderWatchGamePrompt{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvite)] = d.getEventEmitter(func() events.Event {
		return &events.PartySearchInvite{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvites)] = d.getEventEmitter(func() events.Event {
		return &events.PartySearchInvites{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlayerBeaconState)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerBeaconState{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerJoinedChannel{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCOtherLeftChannel)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerLeftChannel{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlaytestStatus)] = d.getEventEmitter(func() events.Event {
		return &events.PlaytestStatus{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCPopup)] = d.getEventEmitter(func() events.Event {
		return &events.Popup{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPrivateCoachingSessionUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.PrivateCoachingSessionUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ProfileCardUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientQuestProgressUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.QuestProgressUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCReadyUpStatus)] = d.getEventEmitter(func() events.Event {
		return &events.ReadyUpStatus{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestLaneSelection)] = d.getEventEmitter(func() events.Event {
		return &events.RequestLaneSelection{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestMMInfo)] = d.getEventEmitter(func() events.Event {
		return &events.RequestMMInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSteamDatagramTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SteamDatagramTicket{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteGCImmediateResponseToInviter{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteReceived{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteResponseReceived{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamsInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTournamentItemDrop)] = d.getEventEmitter(func() events.Event {
		return &events.TournamentItemDrop{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTrophyAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.TrophyAwarded{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientUnderDraftGoldUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.UnderDraftGoldUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientVACReminder)] = d.getEventEmitter(func() events.Event {
		return &events.VACReminder{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientWageringUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.WageringUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchDownloadedReplay)] = d.getEventEmitter(func() events.Event {
		return &events.WatchDownloadedReplay{}
	})
}
