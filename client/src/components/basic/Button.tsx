import React from "react";
import "./Button.css";
import "./Truncated.css";

type Props = {
  value: string;
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
  onClick?: (e: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
};

export default ({ type = "button", ...rest }: Props) => (
  <input type={type} className="Button Truncated" {...rest} />
);
