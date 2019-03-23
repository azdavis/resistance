import React from "react";
import { D } from "../types";
import Button from "./basic/Button";
import ButtonLink from "./basic/ButtonLink";

type Props = {
  d: D;
};

export default ({ d }: Props) => {
  return (
    <div className="Welcome">
      <h1>Resistance</h1>
      <Button value="Play" onClick={() => d({ t: "GoNameChoose" })} />
      <Button value="Learn how to play" onClick={() => d({ t: "GoHowTo" })} />
      <ButtonLink
        value="View source code"
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
