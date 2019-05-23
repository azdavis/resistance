import React, { useState } from "react";
import { Translation } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  reconnect: () => void;
};

export default ({ t, reconnect }: Props) => {
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      <h1>{t.Disconnected.title}</h1>
      <Button
        value={t.Disconnected.reconnect}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
