import React from "react";
import Button from "../basic/Button";

type Props = {};

export default ({  }: Props) => {
  return (
    <div className="Disconnected">
      <h1>Disconnected</h1>
      <Button value="Reconnect" />
    </div>
  );
};
