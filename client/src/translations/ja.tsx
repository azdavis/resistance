import { minN, maxN, maxPts } from "../shared";
import fullWidth from "../fullWidth";

export default {
  resName: "抵抗勢力",
  spyName: "スパイ",
  submit: "送信する",
  leave: "去る",
  back: "戻る",
  Disbanded: {
    title: <h1>解散</h1>,
    body: <p>あなたのいたゲームまたはロビーは解散された。</p>,
  },
  Disconnected: {
    title: <h1>接続が切られた</h1>,
    reconnect: "再接続する",
  },
  Fatal: {
    title: <h1>致命的謝り</h1>,
    body: <p>アプリが復活できぬ謝りが起きた。</p>,
  },
  GamePlaying: {
    viewAllegiance: "忠誠を見る",
    captain: (x: string) => <div>主将：{x}</div>,
    members: (n: number) => <div>使命員（{fullWidth(n)}）：</div>,
    beingChosen: <div>（選択中）</div>,
    succeedPrompt: "使命は成功するか？",
    succeed: "成功",
    fail: "失敗",
    beingVotedOn: <div>（投票中）</div>,
    occurPrompt: "使命は起こるか？",
    occur: "起こる",
    notOccur: "起こらない",
  },
  HowTo: {
    title: <h1>遊び方</h1>,
    groupSize: (
      <p>
        最低{fullWidth(minN)}人、最高{fullWidth(maxN)}人のグループは遊べる。
      </p>
    ),
    groupNames: <p>あるプレイヤーはスパイ。他のプレイヤーは抵抗勢力員。</p>,
    decideWinner: (
      <p>
        スパイと抵抗勢力のどちらかが先に{fullWidth(maxPts)}点を取る方が勝利。
      </p>
    ),
    captain: (
      <p>
        ゲームはラウンドで行う。ラウンドごとに、主将は選ばれる。主将はラウンドの使命員を選ぶ。
      </p>
    ),
    occurVote: (
      <p>
        主将が選び終わった際、プレイヤー全員が使命が起こるかどうか投票する。
      </p>
    ),
    noOccur: <p>使命が起こらなければ、次のラウンドが始まる。</p>,
    tooManyNoOccur: (
      <p>あまりにも多くの使命が連続して起こらなければ、スパイが１点を取る。</p>
    ),
    yesOccur: <p>使命が起これば、使命員が成功するかどうか投票する。</p>,
    succeed: <p>使命が成功すれば、抵抗勢力が１点を取る。</p>,
    fail: <p>使命が失敗すれば、スパイが１点を取る。</p>,
  },
  LangChoosing: {
    title: <h1>言語の設定</h1>,
    langNames: "日本語",
  },
  LobbyWaiting: {
    title: (n: number) => <h1>ロビー（{fullWidth(n)}）</h1>,
    start: "始める",
  },
  NameChoosing: {
    title: <h1>プレイヤー名</h1>,
    invalid: "無効",
  },
  Welcome: {
    play: "遊ぶ",
    learnHow: "遊び方を知る",
    setLang: "言語を設定する",
    viewCode: "コードを見る",
  },
};
