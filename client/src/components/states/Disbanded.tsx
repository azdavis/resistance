import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

export default ({ t, d }: Props) => {
  const { Disbanded: D, leave } = t;
  return (
    <div className="Disbanded">
      <h1>{D.title}</h1>
      <p>{D.body}</p>
      <Button value={leave} onClick={() => d({ t: "GoLobbies" })} />
    </div>
  );
};
