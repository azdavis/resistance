import React, { useState } from "react";
import Button from "./Button";

type Props = {
  isSpy: boolean;
};

export default ({ isSpy }: Props) => {
  const [show, setShow] = useState(false);
  const value = show
    ? `You ${isSpy ? "are" : "are not"} a spy`
    : "View spy status";
  return <Button value={value} onClick={() => setShow(!show)} />;
};
