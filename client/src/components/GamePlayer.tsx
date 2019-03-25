import React from "react";
import { Send, Client, CID } from "../types";
import { getCaptain } from "../consts";
import Voter from "./basic/Voter";
import MemberChooser from "./MemberChooser";

type Props = {
  send: Send;
  me: CID;
  clients: Array<Client>;
  isSpy: boolean;
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

const succeedOpts: Array<[string, boolean]> = [
  ["Succeed", true],
  ["Fail", false],
];

const occurOpts: Array<[string, boolean]> = [
  ["Occur", true],
  ["Not occur", false],
];

export default ({
  send,
  me,
  clients,
  isSpy,
  resPts,
  spyPts,
  captain,
  members,
  active,
}: Props) => (
  <div className="GamePlayer">
    <h1>Game</h1>
    <p>Allegiance: {isSpy ? "Spies" : "Resistance"}</p>
    <p>Resistance points: {resPts}</p>
    <p>Spy points: {spyPts}</p>
    <p>Captain: {getCaptain(clients, captain)}</p>
    <p>Members: ({typeof members === "number" ? members : members.length})</p>
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
          options={succeedOpts}
          onVote={Vote => send({ t: "MissionVote", Vote })}
        />
      ) : (
        <div>(being voted on by mission members)</div>
      )
    ) : (
      <Voter
        prompt="Should the mission occur?"
        options={occurOpts}
        onVote={Vote => send({ t: "MemberVote", Vote })}
      />
    )}
  </div>
);
