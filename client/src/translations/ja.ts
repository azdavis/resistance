import { minN, maxN, maxPts } from "../shared";
import fullWidth from "../fullWidth";

export default {
  langName: "日本語",
  resName: "抵抗勢力",
  spyName: "スパイ",
  submit: "送信する",
  leave: "去る",
  back: "戻る",
  Disbanded: {
    title: "解散",
    body: "あなたのいたゲームまたはロビーは解散された。",
  },
  Disconnected: {
    title: "接続が切られた",
    reconnect: "再接続する",
  },
  Fatal: {
    title: "致命的謝り",
    body: "アプリが復活できぬ謝りが起きた。",
  },
  GamePlaying: {
    viewAllegiance: "忠誠を見る",
    captain: (x: string) => `主将：${x}`,
    members: (n: number) => `使命員（${fullWidth(n)}）：`,
    beingChosen: "（選択中）",
    succeedPrompt: "使命は成功するか？",
    succeed: "成功",
    fail: "失敗",
    beingVotedOn: "（投票中）",
    occurPrompt: "使命は起こるか？",
    occur: "起こる",
    notOccur: "起こらない",
  },
  HowTo: {
    title: "遊び方",
    groupSize:
      "最低" +
      fullWidth(minN) +
      "人、最高" +
      fullWidth(maxN) +
      "人のグループは遊べる。",
    groupNames: "あるプレイヤーはスパイ。他のプレイヤーは抵抗勢力員。",
    decideWinner:
      "スパイと抵抗勢力のどちらかが先に" +
      fullWidth(maxPts) +
      "点を取る方が勝利。",
    captain:
      "ゲームはラウンドで行う。ラウンドごとに、主将は選ばれる。主将はラウンドの使命員を選ぶ。",
    occurVote:
      "主将が選び終わった際、プレイヤー全員が使命が起こるかどうか投票する。",
    noOccur: "使命が起こらなければ、次のラウンドが始まる。",
    tooManyNoOccur:
      "あまりにも多くの使命が連続して起こらなければ、スパイが１点を取る。",
    yesOccur: "使命が起これば、使命員が成功するかどうか投票する。",
    succeed: "使命が成功すれば、抵抗勢力が１点を取る。",
    fail: "使命が失敗すれば、スパイが１点を取る。",
  },
  LangChoosing: {
    title: "言語の設定",
  },
  LobbyChoosing: {
    title: "ロビー",
    create: "新たなのを作成する",
    existing: (n: number) => `存在するロビー（${fullWidth(n)}）`,
  },
  LobbyWaiting: {
    title: (n: number) => `ロビー（${fullWidth(n)}）`,
    start: "始める",
  },
  NameChoosing: {
    title: "プレイヤー名",
    invalid: "無効",
  },
  Welcome: {
    play: "遊ぶ",
    learnHow: "遊び方を知る",
    setLang: "言語を設定する",
    viewCode: "コードを見る",
  },
};
