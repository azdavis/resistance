import React from "react";
import { Lang } from "../../types";
import { maxPts } from "../../shared";
import { resName, spyName } from "../../text";
import Checkbox from "./Checkbox";
import "./Scoreboard.css";

type Props = {
  lang: Lang;
  resPts: number;
  spyPts: number;
};

const points = Array.from(Array(maxPts), (_, i) => i + 1);

export default ({ lang, resPts, spyPts }: Props) => (
  <table className="Scoreboard">
    <tbody>
      <tr>
        <td>
          {points.map(k => (
            <Checkbox key={k} disabled checked={resPts >= k} />
          ))}
        </td>
        <td>
          <b>{resName[lang]}</b>
        </td>
      </tr>
    </tbody>
    <tr>
      <td>
        {points.map(k => (
          <Checkbox key={k} disabled checked={spyPts >= k} />
        ))}
      </td>
      <td>
        <b>{spyName[lang]}</b>
      </td>
    </tr>
  </table>
);
