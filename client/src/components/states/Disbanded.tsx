import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

export default ({ t, d }: Props) => (
  <div className="Disbanded">
    <h1>{t.disbanded}</h1>
    <p>{t.disbandedGameOrLobby}</p>
    <Button value={t.leave} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
