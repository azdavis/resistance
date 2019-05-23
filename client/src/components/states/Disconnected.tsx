import React, { useState } from "react";
import t8ns from "../../translations";
import { Lang } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  reconnect: () => void;
};

export default ({ lang, reconnect }: Props) => {
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      {t8ns[lang].Disconnected.title}
      <Button
        value={t8ns[lang].Disconnected.reconnect}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
