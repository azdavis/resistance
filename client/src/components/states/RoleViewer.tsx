import React from "react";
import { D } from "../../types";
import Button from "../Button";

type Props = {
  d: D;
  isSpy: boolean;
};

export default ({ d, isSpy }: Props) => {
  return (
    <div className="RoleViewer">
      <h1>Role</h1>
      <p>You {isSpy ? "are" : "are not"} a spy.</p>
      <Button value="Continue" onClick={() => d({ t: "AckRole" })} />
    </div>
  );
};
