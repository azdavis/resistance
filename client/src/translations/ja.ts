// spellchecker: disable

import { Translation } from "../etc";
import { minN, maxN, maxPts } from "../shared";
import fullWidth from "../fullWidth";

const ja: Translation = {
  lang: "ja",
  resName: "抵抗勢力",
  spyName: "スパイ",
  submit: "送信する",
  leave: "去る",
  back: "戻る",
  disbanded: "解散",
  disconnected: "接続が切られた",
  errorWithCode: code => `エラー${code}`,
  reconnect: "再接続する",
  invalid: "無効",
  invalidStateTransition: "無効な状態の推移が起きた。",
  viewAllegiance: "忠誠を見る",
  captain: x => `主将：${x}`,
  members: n => `使命員（${fullWidth(n)}）：`,
  beingChosen: "（選択中）",
  succeedPrompt: "使命は成功すべきか？",
  succeed: "成功",
  fail: "失敗",
  beingVotedOn: "（投票中）",
  occurPrompt: "使命は起こるべきか？",
  occur: "起こる",
  notOccur: "起こらない",
  howToPlay: "遊び方",
  setLang: "言語の設定",
  lobbies: "ロビー",
  createNew: "新たなのを作成する",
  existingLobbies: n => `存在するロビー（${fullWidth(n)}）`,
  lobbyWaiting: n => `ロビー（${fullWidth(n)}）`,
  start: "始める",
  playerName: "プレイヤー名",
  play: "遊ぶ",
  learnHow: "遊び方を知る",
  viewCode: "コードを見る",
  howTo: [
    `最低${fullWidth(minN)}人、最高${fullWidth(maxN)}人のグループは遊べる。`,
    "あるプレイヤーはスパイ。他のプレイヤーは抵抗勢力。",
    `スパイと抵抗勢力のどちらかが先に${fullWidth(maxPts)}点を取る方が勝利。`,
    "ゲームはラウンドで行う。ラウンドごとに、主将は選ばれる。主将はラウンドの使命員を選ぶ。",
    "主将が選び終わった際、プレイヤー全員が使命が起こるかどうか投票する。",
    "使命が起こらなければ、次のラウンドが始まる。",
    "あまりにも多くの使命が連続して起こらなければ、スパイが１点を取る。",
    "使命が起これば、使命員が成功するかどうか投票する。",
    "使命が成功すれば、抵抗勢力が１点を取る。",
    "使命が失敗すれば、スパイが１点を取る。",
  ],
};

export default ja;
