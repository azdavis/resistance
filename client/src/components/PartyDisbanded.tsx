import React, { Dispatch } from "react";
import { ToClient, PartyInfo } from "../types";

type Props = {
  parties: Array<PartyInfo>;
  d: Dispatch<ToClient>;
};

export default ({ d, parties }: Props): JSX.Element => {
  return (
    <>
      <h1>Party disbanded</h1>
      <p>The party has been disbanded</p>
      <input
        type="button"
        value="OK"
        onClick={() => d({ T: "PartyChoosing", Parties: parties })}
      />
    </>
  );
};
