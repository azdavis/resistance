import React from "react";
import { D } from "../types";
import { MinN } from "../consts";
import Button from "./Button";

type Props = {
  d: D;
};

export default ({ d }: Props) => (
  <div className="HowTo">
    <h1>How to play</h1>
    <p>Groups of at least {MinN} players may play.</p>
    <p>
      Some of the players are spies. The rest of the players are members of the
      resistance.
    </p>
    <p>The first of the spies and resistance to get 3 points wins the game.</p>
    <p>
      The game is played in rounds. In every round of the game, a mission
      captain is chosen. The mission captain chooses a certain group of players
      to participate in the mission.
    </p>
    <p>
      When the mission captain has finished choosing the members of the mission,
      all players vote on whether the mission occurs. This vote is public. The
      winner is decided by simple majority, with ties resulting in the mission
      not occurring.
    </p>
    <p>If the mission does not occur, the next round is started.</p>
    <p>If too many missions do not occur in a row, the spies get 1 point.</p>
    <p>
      If the mission does occur, the members of the mission vote on whether the
      mission succeeds or not. This vote is private. The winner is decided based
      on how many members were in the mission.
    </p>
    <p>If the mission succeeds, the resistance gets 1 point.</p>
    <p>If the mission fails, the spies get 1 point.</p>
    <Button value="Return" onClick={() => d({ t: "GoNameChoose" })} />
  </div>
);
