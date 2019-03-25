import React from "react";
import { Send, Client, CID } from "../types";
import { getCaptain } from "../consts";
import Voter from "./basic/Voter";
import MemberChooser from "./MemberChooser";

type Props = {
  send: Send;
  me: CID;
  clients: Array<Client>;
  resPts: number;
  spyPts: number;
  captain: CID;
  members: number | Array<CID>;
  active: boolean;
};

const modifiers = (cid: CID, me: CID, captain: CID): string =>
  cid === me && cid === captain
    ? " (you, captain)"
    : cid === me
    ? " (you)"
    : cid === captain
    ? " (captain)"
    : "";

const options: Array<[string, boolean]> = [["Yes", true], ["No", false]];

export default ({
  send,
  me,
  clients,
  resPts,
  spyPts,
  captain,
  members,
  active,
}: Props) => {
  const numMembers = typeof members === "number" ? members : members.length;
  return (
    <div className="GamePlayer">
      <h1>Game</h1>
      <p>Resistance points: {resPts}</p>
      <p>Spy points: {spyPts}</p>
      <p>Captain: {getCaptain(clients, captain)}</p>
      <p>Members ({numMembers}):</p>
      {typeof members === "number" ? (
        me === captain ? (
          <MemberChooser {...{ send, me, clients, members }} />
        ) : (
          <div>(being chosen by captain)</div>
        )
      ) : (
        clients
          .filter(({ CID }) => members.includes(CID))
          .map(({ CID, Name }) => (
            <div key={CID}>
              {Name}
              {modifiers(CID, me, captain)}
            </div>
          ))
      )}
      {typeof members === "number" ? null : active ? (
        members.includes(me) ? (
          <Voter
            prompt="Should the mission succeed?"
            options={options}
            onVote={Vote => send({ t: "MissionVote", Vote })}
          />
        ) : (
          <div>(being voted on by mission members)</div>
        )
      ) : (
        <Voter
          prompt="Should the mission occur?"
          options={options}
          onVote={Vote => send({ t: "MemberVote", Vote })}
        />
      )}
    </div>
  );
};
