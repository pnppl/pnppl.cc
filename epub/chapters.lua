function Header(el)
  if el.classes:includes("level1") then
    -- Keep it at H1 level, don't shift
    return el
  else
    -- Shift other headings down by 1
    el.level = el.level + 1
    return el
  end
end
