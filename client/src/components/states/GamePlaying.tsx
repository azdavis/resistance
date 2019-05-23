import React from "react";
import t8ns from "../../translations";
import { Client, CID } from "../../shared";
import { Lang, S } from "../../etc";
import ButtonSpoiler from "../basic/ButtonSpoiler";
import MemberChooser from "../basic/MemberChooser";
import Scoreboard from "../basic/Scoreboard";
import Voter from "../basic/Voter";
import "../basic/Truncated.css";

type Props = {
  lang: Lang;
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
}: Props) => {
  const t8n = t8ns[lang].GamePlaying;
  return (
    <div className="GamePlaying">
      <Scoreboard {...{ lang, resPts, spyPts }} />
      <ButtonSpoiler
        view={t8n.viewAllegiance}
        spoil={isSpy ? t8ns[lang].spyName : t8ns[lang].resName}
      />
      <div>{t8n.captain(clients.find(({ CID }) => CID === captain)!.Name)}</div>
      <div>{t8n.members(isNum(members) ? members : members.length)}</div>
      {isNum(members) ? (
        me === captain ? (
          <MemberChooser {...{ lang, send, me, clients, members }} />
        ) : (
          t8n.beingChosen
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
            prompt={t8n.succeedPrompt}
            options={[[t8n.succeed, true], [t8n.fail, false]]}
            onVote={Vote => send({ t: "MissionVote", Vote })}
          />
        ) : (
          t8n.beingVotedOn
        )
      ) : (
        <Voter
          // the `key`s must differ
          key="occur"
          prompt={t8n.occurPrompt}
          options={[[t8n.occur, true], [t8n.notOccur, false]]}
          onVote={Vote => send({ t: "MemberVote", Vote })}
        />
      )}
    </div>
  );
};
