import React, { useState } from "react";
import { D } from "../types";
import Button from "./Button";

type Props = {
  d: D;
  isSpy: boolean;
  wait: boolean;
};

export default ({ d, isSpy, wait }: Props) => {
  const [show, setShow] = useState(false);
  const value = show ? `You ${isSpy ? "are" : "are not"} a spy` : "View role";
  return (
    <div className="RoleViewer">
      <h1>Role</h1>
      <Button value={value} onClick={() => setShow(!show)} />
      <Button
        value="Continue"
        onClick={() => d({ t: "AckRole" })}
        disabled={wait}
      />
    </div>
  );
};
