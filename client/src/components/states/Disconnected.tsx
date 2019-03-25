import React, { useState } from "react";
import Button from "../basic/Button";

type Props = {
  reconnect: () => void;
};

export default ({ reconnect }: Props) => {
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      <h1>Disconnected</h1>
      <Button
        value="Reconnect"
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
