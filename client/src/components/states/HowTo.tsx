import React from "react";
import { minN, maxN, maxPts } from "../../shared";
import { Lang, D } from "../../etc";
import { back } from "../../text";
import fullWidth from "../../fullWidth";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>How to play</h1>,
    ja: <h1>遊び方</h1>,
  },
  groupSize: {
    en: (
      <p>
        Groups of at least {minN} and at most {maxN} players may play.
      </p>
    ),
    ja: (
      <p>
        最低{fullWidth(minN)}人、最高{fullWidth(maxN)}人のグループは遊べる。
      </p>
    ),
  },
  groupNames: {
    en: <p>Some players are spies. The rest are members of the resistance.</p>,
    ja: <p>あるプレイヤーはスパイ。他のプレイヤーは抵抗勢力員。</p>,
  },
  decideWinner: {
    en: (
      <p>
        The first of the spies and resistance to get {maxPts} points wins the
        game.
      </p>
    ),
    ja: (
      <p>
        スパイと抵抗勢力のどちらかが先に{fullWidth(maxPts)}点を取る方が勝利。
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
    ja: (
      <p>
        ゲームはラウンドで行う。ラウンドごとに、主将は選ばれる。主将はラウンドの使命員を選ぶ。
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
    ja: (
      <p>
        主将が選び終わった際、プレイヤー全員が使命が起こるかどうか投票する。
      </p>
    ),
  },
  noOccur: {
    en: <p>If the mission does not occur, the next round is started.</p>,
    ja: <p>使命が起こらなければ、次のラウンドが始まる。</p>,
  },
  tooManyNoOccur: {
    en: (
      <p>If too many missions do not occur in a row, the spies get 1 point.</p>
    ),
    ja: (
      <p>あまりにも多くの使命が連続して起こらなければ、スパイが１点を取る。</p>
    ),
  },
  yesOccur: {
    en: (
      <p>
        If the mission does occur, the members of the mission vote on whether
        the mission succeeds.
      </p>
    ),
    ja: <p>使命が起これば、使命員が成功するかどうか投票する。</p>,
  },
  succeed: {
    en: <p>If the mission succeeds, the resistance gets 1 point.</p>,
    ja: <p>使命が成功すれば、抵抗勢力が１点を取る。</p>,
  },
  fail: {
    en: <p>If the mission fails, the spies get 1 point.</p>,
    ja: <p>使命が失敗すれば、スパイが１点を取る。</p>,
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
