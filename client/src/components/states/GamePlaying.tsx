import React from "react";
import { Lang, Send, Client, CID } from "../../types";
import ButtonSpoiler from "../basic/ButtonSpoiler";
import MemberChooser from "../basic/MemberChooser";
import Scoreboard from "../basic/Scoreboard";
import Voter from "../basic/Voter";
import "../basic/Truncated.css";

type Props = {
  lang: Lang;
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

const isNum = (x: any): x is number => typeof x == "number";

export default ({
  lang,
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
    <Scoreboard {...{ lang, resPts, spyPts }} />
    <ButtonSpoiler
      view="View allegiance"
      spoil={isSpy ? "Spies" : "Resistance"}
    />
    <div>Captain: {clients.find(({ CID }) => CID === captain)!.Name}</div>
    <div>Members ({isNum(members) ? members : members.length}):</div>
    {isNum(members) ? (
      me === captain ? (
        <MemberChooser {...{ lang, send, me, clients, members }} />
      ) : (
        <div>(being chosen)</div>
      )
    ) : (
      clients
        .filter(({ CID }) => members.includes(CID))
        .map(({ CID, Name }) => (
          <div key={CID} className="Truncated">
            {Name}
          </div>
        ))
    )}
    {isNum(members) ? null : active ? (
      members.includes(me) ? (
        <Voter
          // the `key`s must differ
          key="succeed"
          prompt="Should the mission succeed?"
          options={[["Succeed", true], ["Fail", false]]}
          onVote={Vote => send({ t: "MissionVote", Vote })}
        />
      ) : (
        <div>(being voted on)</div>
      )
    ) : (
      <Voter
        // the `key`s must differ
        key="occur"
        prompt="Should the mission occur?"
        options={[["Occur", true], ["Not occur", false]]}
        onVote={Vote => send({ t: "MemberVote", Vote })}
      />
    )}
  </div>
);
