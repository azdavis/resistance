import React from "react";
import { Send, Client, CID } from "../../types";
import Voter from "../basic/Voter";
import MemberChooser from "../basic/MemberChooser";

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

const isNum = (x: any): x is number => typeof x == "number";

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
  <div className="GamePlaying">
    <h1>Game</h1>
    <div>Allegiance: {isSpy ? "Spies" : "Resistance"}</div>
    <div>
      Score: Resistance {resPts}, Spies {spyPts}
    </div>
    <div>Captain: {clients.find(({ CID }) => CID === captain)!.Name}</div>
    <div>Members ({isNum(members) ? members : members.length}):</div>
    {isNum(members) ? (
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
    {isNum(members) ? null : active ? (
      members.includes(me) ? (
        <Voter
          // the `key`s must differ
          key="succeed"
          prompt="Should the mission succeed?"
          options={succeedOpts}
          onVote={Vote => send({ t: "MissionVote", Vote })}
        />
      ) : (
        <div>(being voted on by mission members)</div>
      )
    ) : (
      <Voter
        // the `key`s must differ
        key="occur"
        prompt="Should the mission occur?"
        options={occurOpts}
        onVote={Vote => send({ t: "MemberVote", Vote })}
      />
    )}
  </div>
);
