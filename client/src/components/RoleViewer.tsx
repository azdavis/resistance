import React from "react";
import { D } from "../types";
import Button from "./basic/Button";

type Props = {
  d: D;
  isSpy: boolean;
};

export default ({ d, isSpy }: Props) => (
  <div className="RoleViewer">
    <h1>Role</h1>
    <p>You are a {isSpy ? "spy" : "resistance member"}.</p>
    <Button value="Ok" onClick={() => d({ t: "AckRole" })} />
  </div>
);
