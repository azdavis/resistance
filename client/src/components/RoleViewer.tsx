import React from "react";
import { D } from "../types";
import SpyStatus from "./SpyStatus";
import Button from "./Button";

type Props = {
  d: D;
  isSpy: boolean;
};

export default ({ d, isSpy }: Props) => (
  <div className="RoleViewer">
    <h1>Role</h1>
    <SpyStatus isSpy={isSpy} />
    <Button value="Continue" onClick={() => d({ t: "AckRole" })} />
  </div>
);
