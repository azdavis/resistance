import React from "react";
import { maxPts } from "../../shared";
import { Translation } from "../../etc";
import Checkbox from "./Checkbox";
import "./Scoreboard.css";

type Props = {
  t: Translation;
  resPts: number;
  spyPts: number;
};

const points = Array.from(Array(maxPts), (_, i) => i + 1);

export default ({ t, resPts, spyPts }: Props) => (
  <table className="Scoreboard">
    <tbody>
      <tr>
        <td>
          {points.map(k => (
            <Checkbox key={k} disabled checked={resPts >= k} />
          ))}
        </td>
        <td>
          <b>{t.resName}</b>
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
        <b>{t.spyName}</b>
      </td>
    </tr>
  </table>
);
