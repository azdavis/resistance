import React from "react";
import { Lang, Send, Client, CID } from "../../types";
import { resName, spyName } from "../../text";
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

const text = {
  viewAllegiance: {
    en: "View allegiance",
  },
  captain: {
    en: (x: string) => <div>Captain: {x}</div>,
  },
  members: {
    en: (n: number) => <div>Members ({n}):</div>,
  },
  beingChosen: {
    en: <div>(being chosen)</div>,
  },
  succeedPrompt: {
    en: "Should the mission succeed?",
  },
  succeed: {
    en: "Succeed",
  },
  fail: {
    en: "Fail",
  },
  beingVotedOn: {
    en: <div>(being voted on)</div>,
  },
  occurPrompt: {
    en: "Should the mission occur?",
  },
  occur: {
    en: "Occur",
  },
  notOccur: {
    en: "Not occur",
  },
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
      view={text.viewAllegiance[lang]}
      spoil={(isSpy ? resName : spyName)[lang]}
    />
    {text.captain[lang](clients.find(({ CID }) => CID === captain)!.Name)}
    {text.members[lang](isNum(members) ? members : members.length)}
    {isNum(members) ? (
      me === captain ? (
        <MemberChooser {...{ lang, send, me, clients, members }} />
      ) : (
        text.beingChosen[lang]
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
          prompt={text.succeedPrompt[lang]}
          options={[[text.succeed[lang], true], [text.fail[lang], false]]}
          onVote={Vote => send({ t: "MissionVote", Vote })}
        />
      ) : (
        text.beingVotedOn[lang]
      )
    ) : (
      <Voter
        // the `key`s must differ
        key="occur"
        prompt={text.occurPrompt[lang]}
        options={[[text.occur[lang], true], [text.notOccur[lang], false]]}
        onVote={Vote => send({ t: "MemberVote", Vote })}
      />
    )}
  </div>
);
