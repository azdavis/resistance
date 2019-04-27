import React from "react";
import { Client, CID } from "../../shared";
import { Lang, Send } from "../../etc";
import { resName, spyName } from "../../text";
import fullWidth from "../../fullWidth";
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
    ja: "忠誠を見る",
  },
  captain: {
    en: (x: string) => <div>Captain: {x}</div>,
    ja: (x: string) => <div>主将：{x}</div>,
  },
  members: {
    en: (n: number) => <div>Members ({n}):</div>,
    ja: (n: number) => <div>使命員（{fullWidth(n)}）：</div>,
  },
  beingChosen: {
    en: <div>(being chosen)</div>,
    ja: <div>（選択中）</div>,
  },
  succeedPrompt: {
    en: "Should the mission succeed?",
    ja: "使命は成功するか？",
  },
  succeed: {
    en: "Succeed",
    ja: "成功",
  },
  fail: {
    en: "Fail",
    ja: "失敗",
  },
  beingVotedOn: {
    en: <div>(being voted on)</div>,
    ja: <div>（投票中）</div>,
  },
  occurPrompt: {
    en: "Should the mission occur?",
    ja: "使命は起こるか？",
  },
  occur: {
    en: "Occur",
    ja: "起こる",
  },
  notOccur: {
    en: "Not occur",
    ja: "起こらない",
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
      spoil={(isSpy ? spyName : resName)[lang]}
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
