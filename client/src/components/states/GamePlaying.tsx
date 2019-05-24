import React from "react";
import { Client, CID } from "../../shared";
import { Translation, S } from "../../etc";
import ButtonSpoiler from "../basic/ButtonSpoiler";
import MemberChooser from "../basic/MemberChooser";
import Scoreboard from "../basic/Scoreboard";
import Voter from "../basic/Voter";
import "../basic/Truncated.css";

type Props = {
  t: Translation;
  send: S;
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
  t,
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
    <Scoreboard {...{ t, resPts, spyPts }} />
    <ButtonSpoiler
      view={t.viewAllegiance}
      spoil={isSpy ? t.spyName : t.resName}
    />
    <div>{t.captain(clients.find(({ CID }) => CID === captain)!.Name)}</div>
    <div>{t.members(isNum(members) ? members : members.length)}</div>
    {isNum(members) ? (
      me === captain ? (
        <MemberChooser {...{ t, send, me, clients, members }} />
      ) : (
        t.beingChosen
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
          prompt={t.succeedPrompt}
          options={[[t.succeed, true], [t.fail, false]]}
          onVote={Vote => send({ t: "MissionVote", Vote })}
        />
      ) : (
        t.beingVotedOn
      )
    ) : (
      <Voter
        // the `key`s must differ
        key="occur"
        prompt={t.occurPrompt}
        options={[[t.occur, true], [t.notOccur, false]]}
        onVote={Vote => send({ t: "MemberVote", Vote })}
      />
    )}
  </div>
);
