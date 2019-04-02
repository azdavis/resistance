import React from "react";
import "./Button.css";
import "./Truncated.css";
import "./Toggle.css";

type Props = {
  value: string;
  checked?: boolean;
  disabled?: boolean;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

export default ({ value, ...rest }: Props) => (
  <label className="Button Toggle">
    <input type="checkbox" {...rest} />
    <div className="Truncated">{value}</div>
  </label>
);
