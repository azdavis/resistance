import React from "react";
import { D } from "../../types";
import { minN, maxN, maxPts } from "../../consts";
import Button from "../basic/Button";

type Props = {
  d: D;
};

export default ({ d }: Props) => (
  <div className="HowTo">
    <h1>How to play</h1>
    <p>
      Groups of at least {minN} and at most {maxN} players may play.
    </p>
    <p>Some players are spies. The rest are members of the resistance.</p>
    <p>
      The first of the spies and resistance to get {maxPts} points wins the
      game.
    </p>
    <p>
      The game is played in rounds. In every round of the game, a captain is
      chosen. The captain chooses some players to participate in a mission.
    </p>
    <p>
      When the captain has finished choosing, all players vote on whether the
      mission occurs.
    </p>
    <p>If the mission does not occur, the next round is started.</p>
    <p>If too many missions do not occur in a row, the spies get 1 point.</p>
    <p>
      If the mission does occur, the members of the mission vote on whether the
      mission succeeds.
    </p>
    <p>If the mission succeeds, the resistance gets 1 point.</p>
    <p>If the mission fails, the spies get 1 point.</p>
    <Button value="Back" onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
