export interface TreeNode {
  data: Record<string, any>;
  children?: TreeNode[];
}

/**
 * 将树结构转换为一维数组，每个节点包含层级信息。
 * @param soruce 树结构的节点数组，每个节点应包含数据和子节点信息。
 * @param level 当前处理节点的层级，默认为0，表示根节点层级。
 * @returns 返回转换后的一维数组，每个元素包含节点数据和层级信息。
 */
export function treeToLine(soruce: TreeNode[], level: number = 0): any[] {
  const result: any[] = [];

  soruce.forEach((el: any) => {
    result.push({
      _level: level,
      ...el.node,
    })
    if (Array.isArray(el.children)) {
      result.push(...treeToLine(el.children, level + 1));
    }
  })
  return result;

}