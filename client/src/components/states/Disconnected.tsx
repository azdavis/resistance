import React, { useState } from "react";
import { Translation } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  code: number;
  reconnect: () => void;
};

export default ({ t, code, reconnect }: Props) => {
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      <h1>{t.disconnected}</h1>
      <p>{t.errorWithCode(code)}</p>
      <Button
        value={t.reconnect}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
