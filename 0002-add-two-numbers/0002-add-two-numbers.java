/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ArrayList<Integer> arr = new ArrayList<>();
        while ((l1 != null) || (l2 != null)) {
            int val = 0;
            if (l1 != null) {
                val += l1.val;
                l1 = l1.next;
            }

            if (l2 != null) {
                val += l2.val;
                l2 = l2.next;
            }

            arr.add(val%10);
            if ((val / 10) > 0) {
                if (l2 == null) l2 = new ListNode(0);
                l2.val++;
            }
        }

        ListNode res = new ListNode();
        ListNode ptr = res;
        for (int i =0; i <arr.size() ; i++) {
            ptr.val = arr.get(i);
            if (i != arr.size()-1) {
                ptr.next = new ListNode();
                ptr = ptr.next;
            }
        }

        return res;
    }
}