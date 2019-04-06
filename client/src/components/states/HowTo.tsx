import React from "react";
import { Lang, D } from "../../types";
import { minN, maxN, maxPts } from "../../shared";
import { back } from "../../text";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>How to play</h1>,
  },
  groupSize: {
    en: (
      <p>
        Groups of at least {minN} and at most {maxN} players may play.
      </p>
    ),
  },
  groupNames: {
    en: <p>Some players are spies. The rest are members of the resistance.</p>,
  },
  decideWinner: {
    en: (
      <p>
        The first of the spies and resistance to get {maxPts} points wins the
        game.
      </p>
    ),
  },
  captain: {
    en: (
      <p>
        The game is played in rounds. In every round of the game, a captain is
        chosen. The captain chooses the mission members for this round.
      </p>
    ),
  },
  occurVote: {
    en: (
      <p>
        When the captain has finished choosing, all players vote on whether the
        mission occurs.
      </p>
    ),
  },
  noOccur: {
    en: <p>If the mission does not occur, the next round is started.</p>,
  },
  tooManyNoOccur: {
    en: (
      <p>If too many missions do not occur in a row, the spies get 1 point.</p>
    ),
  },
  yesOccur: {
    en: (
      <p>
        If the mission does occur, the members of the mission vote on whether
        the mission succeeds.
      </p>
    ),
  },
  succeed: {
    en: <p>If the mission succeeds, the resistance gets 1 point.</p>,
  },
  fail: {
    en: <p>If the mission fails, the spies get 1 point.</p>,
  },
};

export default ({ lang, d }: Props) => (
  <div className="HowTo">
    {text.title[lang]}
    {text.groupSize[lang]}
    {text.groupNames[lang]}
    {text.decideWinner[lang]}
    {text.captain[lang]}
    {text.occurVote[lang]}
    {text.noOccur[lang]}
    {text.tooManyNoOccur[lang]}
    {text.yesOccur[lang]}
    {text.succeed[lang]}
    {text.fail[lang]}
    <Button value={back[lang]} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
