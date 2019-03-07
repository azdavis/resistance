import React from "react";
import "./Button.css";
import "./FullWidth.css";

type Props = {
  value: string;
  submit?: boolean;
  disabled?: boolean;
  onClick?: (event: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
};

export default ({ submit, ...rest }: Props) => (
  <input
    type={submit ? "submit" : "button"}
    className="Button FullWidth"
    {...rest}
  />
);
