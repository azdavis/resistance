import React, { useState } from "react";
import { Translation } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  reconnect: () => void;
};

export default ({ t, reconnect }: Props) => {
  const D = t.Disconnected;
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      <h1>{D.title}</h1>
      <Button
        value={D.reconnect}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
