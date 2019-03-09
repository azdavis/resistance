import React from "react";
import "./Button.css";
import "./Truncated.css";
import "./Toggle.css";

type Props = {
  value: string;
  checked?: boolean;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

export default ({ value, checked, onChange }: Props) => (
  <label className="Button Toggle">
    <input type="checkbox" checked={checked} onChange={onChange} />
    <div className="Truncated">{value}</div>
  </label>
);
