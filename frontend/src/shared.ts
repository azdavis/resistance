// This should be kept in sync with /server/shared.go.

export const minN = 5;
export const maxN = 7;
export const okGameSize = (n: number): boolean => minN <= n && n <= maxN;
export const maxPts = 3;
export const maxSkip = 3;

export type CID = number;
export type GID = number;

export type Lobby = { GID: GID; Leader: string };
export type Client = { CID: CID; Name: string };

export type ToServer =
  | { t: "Connect" }
  | { t: "Reconnect"; Me: CID; GID: GID }
  | { t: "NameChoose"; Name: string }
  | { t: "LobbyChoose"; GID: GID }
  | { t: "LobbyLeave" }
  | { t: "LobbyCreate" }
  | { t: "GameStart" }
  | { t: "MemberChoose"; Members: Array<CID> }
  | { t: "MemberVote"; Vote: boolean }
  | { t: "MissionVote"; Vote: boolean }
  | { t: "GameLeave" };

export type CurrentGame = {
  t: "CurrentGame";
  IsSpy: boolean;
  ResPts: number;
  SpyPts: number;
  Captain: CID;
  NumMembers: number;
  Members: Array<CID> | null;
  Active: boolean;
};

export type ToClient =
  | { t: "SetMe"; Me: CID }
  | { t: "NameReject" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; GID: GID; Leader: CID; Clients: Array<Client> }
  | CurrentGame
  | {
      t: "EndGame";
      ResPts: number;
      SpyPts: number;
      Lobbies: Array<Lobby>;
    };
