import React, { Dispatch } from "react";
import { ToClient } from "../types";

type Props = {
  d: Dispatch<ToClient>;
};

export default ({ d }: Props): JSX.Element => {
  return (
    <>
      <h1>Party disbanded</h1>
      <p>The party has been disbanded</p>
      <input
        type="button"
        value="OK"
        onClick={() => d({ T: "PartyChoosing" })}
      />
    </>
  );
};
