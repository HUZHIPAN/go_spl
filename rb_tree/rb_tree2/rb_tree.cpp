#include <iostream>
#include <queue>

enum Color {
	BLACK,
	RED
};

struct Node
{
	int val_;
	Node* left_;
	Node* right_;
	Node* parent_;
	Color color_;
	Node(int val = INT_MAX, Node* left = nullptr, Node* right = nullptr,
		Node* parent = nullptr, Color color = BLACK) :
		val_(val), left_(left), right_(right), parent_(parent), color_(color) {
	};
};

class RB_tree {	

public:
	RB_tree() : root_(nullptr) {}

	// 返回节点的颜色
	Color color(Node* node) {
		return nullptr == node ? BLACK : node->color_;
	}
	// 返回节点的左孩子
	Node* left(Node* node) {
		return node->left_;
	}
	// 返回节点的右孩子
	Node* right(Node* node) {
		return node->right_;
	}
	// 返回节点的父节点
	Node* parent(Node* node) {
		return node->parent_;
	}
	// 设置颜色
	void setColor(Node* node, Color color) {
		node->color_ = color;
	}
	// 左旋
	void leftRoate(Node* node) {
		Node* child = node->right_;
		child->parent_ = node->parent_;
		if (node->parent_ == nullptr) { 
			// node本身是根节点 子节点转上去就会变成根节点
			root_ = child;
		} else {
			if (node->parent_->left_ == node) {
				// node是父节点的左孩子
				node->parent_->left_ = child;
			} else {
				// node是父节点的右孩子
				node->parent_->right_ = child;
			}
		}

		node->right_ = child->left_;
		if (child->left_ != nullptr) {
			child->left_->parent_ = node;
		}

		child->left_ = node;
		node->parent_ = child;
	}
	// 右旋操作
	void rightRoate(Node* node) {
		Node* child = node->left_;
		child->parent_ = node->parent_;
		if (node->parent_ == nullptr) {
			// node是根节点
			root_ = child;
		} else {
			if (node->parent_->left_ == node) {
				node->parent_->left_ = child;
			} else {
				node->parent_->right_ = child;
			}
		}

		node->left_ = child->right_;
		if (child->right_ != nullptr) {
			child->right_->parent_ = node;
		}

		child->right_ = node;
		node->parent_ = child;
	}
	void insert(const int& val) {
		// 如果为空树的话 new一个结点出来即可
		if (root_ == nullptr) {
			root_ = new Node(val);
			return;
		}
		//std::cout << val << std::endl;

		Node* parent = nullptr;
		Node* cur = root_;
		while (cur != nullptr) {
			if (cur->val_ > val) {
				parent = cur;
				cur = cur->left_;
			} else if (cur->val_ < val) {
				parent = cur;
				cur = cur->right_;
			} else {
				// 树里面已经有该值了 直接返回即可
				return;
			}
		}

		Node* newNode = new Node(val, nullptr, nullptr, parent, RED);
		if (parent->val_ > val) {
			parent->left_ = newNode;
		} else {
			parent->right_ = newNode;
		}
		// 插入的节点一定是红色
		// 如果父节点的颜色是红色，那么在插入之后需要进行调整
		if (RED == color(parent)) {
			fixAfterInsert(newNode);
		}
	}

	void fixAfterInsert(Node* node) {
		// 判断是否需要迭代检查
		while (RED == color(parent(node))) {
			// 如果插入节点node的父亲是爷爷节点的左孩子
			if (left(parent(parent(node))) == node) {
				// 那么叔叔节点就是爷爷节点的右孩子
				Node* uncle = right(parent(parent(node)));
				if (RED == color(uncle)) { // 情况1和情况3：叔叔节点是红色的
					setColor(parent(node), BLACK);
					setColor(parent(parent(node)), RED);
					setColor(uncle, BLACK);
					node = parent(parent(node)); // 检查节点从插入节点->爷爷节点（因为爷爷的父亲可能是红色）
				} else {
					// 先判断是不是情况三 如果是的话 把情况三调整为情况2 再后续处理
					if (right(parent(node)) == node) {
						leftRoate(parent(node)); // 对插入节点node的父左旋，掰到同一直线
					}

					// 处理情况2
					setColor(parent(node), BLACK);
					setColor(parent(parent(node)), RED);
					rightRoate(parent(parent(node)));
				}
			} else { // 如果插入节点node的父亲是爷爷节点的右孩子,处理完全镜像，所有left改成right即可
				// 那么叔叔节点就是爷爷节点的右孩子
				Node* uncle = left(parent(parent(node)));
				if (RED == color(uncle)) { // 情况1和情况3：叔叔节点是红色的
					setColor(parent(node), BLACK);
					setColor(parent(parent(node)), RED);
					setColor(uncle, BLACK);
					node = parent(parent(node)); // 检查节点从插入节点->爷爷节点（因为爷爷的父亲可能是红色）
				} else {
					// 先判断是不是情况三 如果是的话 把情况三调整为情况2 再后续处理
					if (left(parent(node)) == node) {
						rightRoate(parent(node)); // 对插入节点node的父左旋，掰到同一直线
					}

					// 处理情况2
					setColor(parent(node), BLACK);
					setColor(parent(parent(node)), RED);
					leftRoate(parent(parent(node)));
				}
			}
		}

		setColor(root_, BLACK); // 强制设置根节点为黑色
	}
	
	// 删除
	void remove(const int& val) {
		if (root_ == nullptr) {
			return;
		}

		Node* cur = root_;
		while (nullptr != cur) {
			if (cur->val_ > val) {
				cur = cur->left_;
			} else if (cur->val_ < val) {
				cur = cur->right_;
			} else {
				break;
			}
		}

		if (nullptr == cur) {
			// 没找到 直接返回
			return;
		} 

		// 1.如果待删除的节点有左右孩子的话，我们就把前驱节点覆盖到待删除节点
		// 然后去删除前驱节点
		if (nullptr != cur->left_ && nullptr != cur->right_) {
			Node* pre = cur->left_;
			while (nullptr != pre->right_) {
				pre = pre->right_;
			}
			// 在下面删除前驱节点
			cur->val_ = pre->val_;
			cur = pre;
		}

		// 2.到这里的话 最多只有一个孩子
		Node* child = cur->left_;
		if (cur->right_) {
			child = cur->right_;
		}
		// 如果确实有一个孩子
		if (nullptr != child) {
			child->parent_ = cur->parent_;
			if (cur->parent_ == nullptr) {
				// 说明待删除节点是根节点
				root_ = child;
				/*setColor(root_, BLACK);
				return;*/
			} else {
				if (cur->parent_->left_ == cur) {
					// 说明cur是其父亲的左孩子
					cur->parent_->left_ = child;
				} else {
					// 说明cur是其父亲的右孩子
					cur->parent_->right_ = child;
				}
			}

			Color deleteColor = cur->color_;
			delete cur;
			// 删除的节点是黑色节点才要检查
			if (BLACK == deleteColor) {
				fixAfterRemove(child);
			}
		} else {
			// child = nullptr
			// 删除的节点是叶子节点
			if (nullptr == cur->parent_) {
				// 而且是根节点
				delete cur;
				root_ = nullptr;
				return;
			} else {
				if (color(cur) == BLACK) {
					// 先调整，再删除 ???
					fixAfterRemove(cur);
				} 

				if (cur->parent_->left_ == cur) {
					cur->parent_->left_ = nullptr;
				} else {
					cur->parent_->right_ = nullptr;
				}
				delete cur;
			}
		}
	}

	// 开始检查
	void fixAfterRemove(Node* node) {
		while (node != root_ && color(node) == BLACK) {
			if (left(parent(node)) == node) {
				// 待删除的黑色节点在左子树
				Node* brother = right(parent(node));
				if (color(brother) == RED) {
					// 情况4
					setColor(parent(node), RED);	// 把父亲染红
					setColor(brother, BLACK);		// 把兄弟染黑
					leftRoate(parent(node));		// 左旋父节点
					brother = right(parent(node));	// 重置兄弟
				} 
				if (color(left(brother)) == BLACK && color(right(brother)) == BLACK) {
					// 情况3  
					setColor(brother, RED); // 拉兄弟下水，染兄弟为红色
					node = parent(node);	// 迭代检查
				} else {
					// 情况1、2 进入这里 执行完就会退出
					if (color(left(brother)) == RED && color(right(brother)) == BLACK) {
						// 情况2
				/*		rightRoate(brother);
						brother = right(parent(node));
						setColor(brother, BLACK);
						setColor(right(brother), RED);*/

						setColor(brother, RED);
						setColor(left(brother), BLACK);
						rightRoate(brother);
						brother = right(parent(node));
					}

					//情况1
					setColor(brother, color(parent(node)));
					setColor(parent(node), BLACK);
					setColor(right(brother), BLACK);
					leftRoate(parent(node));
					break;
				}
			} else {
				// 待删除的黑色节点在右子树 把上面涉及方向的操作都反一下就行
 				Node* brother = left(parent(node));
				if (color(brother) == RED) {
					// 情况4
					setColor(parent(node), RED);	// 把父亲染红
					setColor(brother, BLACK);		// 把兄弟染黑
					rightRoate(parent(node));		// 左旋父节点
					brother = left(parent(node));	// 重置兄弟
				}
				if (color(left(brother)) == BLACK && color(right(brother)) == BLACK) {
					// 情况3  
					setColor(brother, RED); // 拉兄弟下水，染兄弟为红色
					node = parent(node);	// 迭代检查
				} else {
					// 情况1、2 进入这里 执行完就会退出
					if (color(left(brother)) == RED && color(right(brother)) == BLACK) {
						// 情况2
						/*leftRoate(brother);
						brother = left(parent(node));
						setColor(brother, BLACK);
						setColor(left(brother), RED);*/

						setColor(brother, RED);
						setColor(right(brother), BLACK);
						leftRoate(brother);
						brother = left(parent(node));
					}

					//情况1
					setColor(brother, color(parent(node)));
					setColor(parent(node), BLACK);
					setColor(left(brother), BLACK);
					rightRoate(parent(node));
					
					break;
				}
			}
		}

		// 如果发现node指向的节点是红色，直接涂成黑色，调整结束
		setColor(node, BLACK);
	}

	// 返回根节点
	Node* getRoot() {
		return root_;
	}

private:
	Node* root_;
};


void printTree(Node* root) {
	std::queue<Node*> que;
	if (root) que.push(root);

	while (!que.empty()) {
		int size = que.size();

		for (int i = 0; i < size; ++i) {
			Node* cur = que.front(); que.pop();
			if (cur != nullptr) {
				std::cout << cur->val_ << "(" << (cur->color_ == 0 ? "B" : "R") << ")" << " ";
				que.push(cur->left_);
				que.push(cur->right_);
			} else {
				std::cout << "NULL ";
			}
		}
		std::cout << std::endl;
	}
}

int main() {
	
	{   // 测试删除操作
		std::cout << "==========插入============" << std::endl;
		RB_tree rbTree;
		for (int i = 1; i <= 10; ++i) {
			rbTree.insert(i);
		}
		printTree(rbTree.getRoot());

		std::cout << "==========删除============" << std::endl;
		rbTree.remove(9);
		rbTree.remove(10);
		printTree(rbTree.getRoot());
	}
	return 0;
}
